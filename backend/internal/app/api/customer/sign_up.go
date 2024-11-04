package customer

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/render"

	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
)

// CustomerSignUp godoc
// @Summary      SingUp customer
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        input body       model.CustomerAliasEmailPasswordRequest true "sign up"
// @Success      200   {object}   model.ErrorResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      409   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router       /customer/sign-up [post]
func (i *Customer) CustomerSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CustomerAliasEmailPasswordRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[CustomerSignUp] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CustomerSignUp] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	id, err := i.srv.CreateCustomer(ctx, req.Alias, req.Email, req.Password)
	if errors.Is(err, model.ErrInvalidEmail) {
		logger.Error(ctx, "[CustomerSignUp] email is invalid", "email", req.Email)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("email is invalid"))
		return
	}
	if errors.Is(err, model.ErrEmailRegistered) {
		logger.Error(ctx, "[CustomerSignUp] email is already taken", "email", req.Email)
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, model.NewErrResp("email is already taken"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CustomerSignUp] failed to create customer", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to create customer"))
		return
	}

	token, err := middleware.GenerateToken(id, middleware.RoleCustomer, req.Alias, i.srv.GetJWTSecret())
	if err != nil {
		logger.Error(ctx, "[CustomerSignUp] failed to generate token", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to generate token"))
		return
	}

	logger.Info(ctx, "[CustomerSignUp] customer created", "alias", req.Alias, "email", req.Email)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &model.CustomerTokenResponse{
		Token: token,
	})
}
