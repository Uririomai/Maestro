package order

import (
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"net/http"
)

// GetMyOrders godoc
// @Summary Get customers orders
// @Security ApiKeyAuth
// @Tags order
// @Accept   json
// @Produce  json
// @Success      200   {object}   model.OrderDTO
// @Failure      403   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /order/get-my [get]
func (i *Order) GetMyOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsCustomer(ctx) {
		logger.Error(ctx, "[GetMyOrders] user is not customer")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	customerId := middleware.GetUserId(ctx)

	orders, err := i.srv.GetOrdersByCustomerId(ctx, customerId)
	if err != nil {
		logger.Error(ctx, "[GetMyOrders] failed to get orders", "customer_id", customerId, "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get orders"))
		return
	}

	logger.Info(ctx, "[GetMyOrders] orders giver", "customer_id", customerId)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromOrdersToDTO(orders))
}
