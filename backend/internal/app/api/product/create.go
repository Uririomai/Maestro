package product

import (
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// CreateProduct godoc
// @Summary Create product
// @Security ApiKeyAuth
// @Tags product
// @Accept   json
// @Produce  json
// @Param input body model.CreateProductRequest true "product info"
// @Success      200   {object}   model.ProductDTO
// @Failure      400   {object}   model.ErrorResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      404   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /product/create [post]
func (i *Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsAdmin(ctx) {
		logger.Error(ctx, "[CreateProduct] user is not admin")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	adminId := middleware.GetUserId(ctx)

	var req model.CreateProductRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[CreateProduct] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CreateProduct] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	website, err := i.srv.GetWebsiteByAlias(ctx, req.WebsiteAlias)
	if errors.Is(err, model.ErrWebsiteNotFound) {
		logger.Error(ctx, "[CreateProduct] website not found", "alias", req.WebsiteAlias)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, model.NewErrResp("website not found"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CreateProduct] failed to check admin", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to check admin"))
		return
	}
	if website.AdminId != adminId {
		logger.Error(ctx, "[CreateProduct] admin is not owner")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	product, err := i.srv.CreateProduct(ctx, model.FromCreateRequestToProduct(&req))
	if err != nil {
		logger.Error(ctx, "[CreateProduct] failed to create product", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to create product"))
		return
	}

	logger.Info(ctx, "[CreateProduct] product created", "id", product.Id)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromProductToDTO(product))
}
