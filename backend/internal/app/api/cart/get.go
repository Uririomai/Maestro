package cart

import (
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"net/http"
)

// GetCart godoc
// @Summary Get customer's cart
// @Security ApiKeyAuth
// @Tags cart
// @Accept   json
// @Produce  json
// @Success      200   {object}   model.CartResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /cart/get [get]
func (i *Cart) GetCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsCustomer(ctx) {
		logger.Error(ctx, "[GetCart] user is not customer")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	customerId := middleware.GetUserId(ctx)

	cart, err := i.srv.GetCart(ctx, customerId)
	if err != nil {
		logger.Error(ctx, "[GetCart] failed to get cart", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get cart"))
		return
	}

	logger.Info(ctx, "[GetCart] cart given", "customer_id", customerId)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromCartToDTO(cart))
}
