package cart

import (
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// AddProductToCart godoc
// @Summary Add product to customer's cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept   json
// @Produce  json
// @Param input body model.AddProductToCartRequest true "product id and count"
// @Success      200   {object}   model.CartResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      404   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /cart/add-product [post]
func (i *Cart) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsCustomer(ctx) {
		logger.Error(ctx, "[AddProductToCart] user is not customer")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	customerId := middleware.GetUserId(ctx)

	var req model.AddProductToCartRequest
	// TODO: валидация отрицательного кол-ва
	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[AddProductToCart] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[AddProductToCart] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	product, err := i.srv.GetProductById(ctx, req.ProductId)
	if errors.Is(err, model.ErrNotFound) {
		logger.Error(ctx, "[AddProductToCart] product not found", "id", req.ProductId)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, model.NewErrResp("product not found"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[AddProductToCart] failed to check product", "id", req.ProductId, "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to add product"))
		return
	}
	if product.WebsiteAlias != middleware.GetWebsiteAlias(ctx) {
		logger.Error(ctx, "[AddProductToCart] product is not from this site")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("product is not from this site"))
		return
	}

	cart, err := i.srv.AddProductToCart(ctx, customerId, req.ProductId, req.Count)
	if err != nil {
		logger.Error(ctx, "[AddProductToCart] failed to add product to cart", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to add product to cart"))
		return
	}

	logger.Info(ctx, "[AddProductToCart] product added to cart", "customer_id", customerId)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromCartToDTO(cart))
}
