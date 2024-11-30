package order

import (
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// MakeOrder godoc
// @Summary Make order by cart
// @Security ApiKeyAuth
// @Tags order
// @Accept   json
// @Produce  json
// @Param input body model.MakeOrderRequest true "customer comment to order"
// @Success      200   {object}   model.OrderDTO
// @Failure      400   {object}   model.ErrorResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /order/make [post]
func (i *Order) MakeOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsCustomer(ctx) {
		logger.Error(ctx, "[MakeOrder] user is not customer")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	customerId := middleware.GetUserId(ctx)

	var req model.MakeOrderRequest
	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[MakeOrder] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[MakeOrder] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	order, err := i.srv.CreateOrder(ctx, customerId, req.Comment)
	if errors.Is(err, model.ErrEmptyOrder) {
		logger.Error(ctx, "[MakeOrder] order is empty", "customer_id", customerId)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("order is empty"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[MakeOrder] failed to make order", "customer_id", customerId, "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to make order"))
		return
	}

	logger.Info(ctx, "[MakeOrder] order created", "order_id", order.Id)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromOrderToDTO(order))
}
