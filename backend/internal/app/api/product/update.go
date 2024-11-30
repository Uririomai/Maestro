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

// UpdateProduct godoc
// @Summary Update product
// @Security ApiKeyAuth
// @Tags product
// @Accept   json
// @Produce  json
// @Param input body model.UpdateProductRequest true "product info"
// @Success      200   {object}   model.ProductDTO
// @Failure      400   {object}   model.ErrorResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      404   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /product/update [put]
func (i *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsAdmin(ctx) {
		logger.Error(ctx, "[UpdateProduct] user is not admin")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	adminId := middleware.GetUserId(ctx)

	var req model.UpdateProductRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[UpdateProduct] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[UpdateProduct] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	product, err := i.srv.GetProductById(ctx, req.Id)
	if errors.Is(err, model.ErrNotFound) {
		logger.Error(ctx, "[UpdateProduct] product not found", "id", req.Id)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, model.NewErrResp("product not found"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[UpdateProduct] failed to check product", "id", req.Id, "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to check product"))
		return
	}

	website, err := i.srv.GetWebsiteByAlias(ctx, product.WebsiteAlias)
	if errors.Is(err, model.ErrNotFound) {
		logger.Error(ctx, "[UpdateProduct] website not found", "alias", product.WebsiteAlias)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, model.NewErrResp("website not found"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[UpdateProduct] failed to check admin", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to check admin"))
		return
	}
	if website.AdminId != adminId {
		logger.Error(ctx, "[UpdateProduct] admin is not owner")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	product, err = i.srv.UpdateProduct(ctx, model.FromUpdateRequestToProduct(&req))
	if err != nil {
		logger.Error(ctx, "[UpdateProduct] failed to update product", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to update product"))
		return
	}

	logger.Info(ctx, "[UpdateProduct] product updated", "id", product.Id)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromProductToDTO(product))
}
