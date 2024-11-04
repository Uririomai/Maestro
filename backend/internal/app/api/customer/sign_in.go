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

// CustomerSignIn godoc
// @Summary      SingIn customer
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        input body       model.CustomerAliasEmailPasswordRequest true "sign in"
// @Success      200   {object}   model.ErrorResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router       /customer/sign-in [post]
func (i *Customer) CustomerSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CustomerAliasEmailPasswordRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[CustomerSignIn] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CustomerSignIn] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	// TODO: мб добавить WebsiteIsExist

	admin, err := i.srv.GetCustomerIdByEmailPassword(ctx, req.Alias, req.Email, req.Password)
	if errors.Is(err, model.ErrInvalidEmail) {
		logger.Error(ctx, "[CustomerSignIn] email is invalid", "email", req.Email)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("email is invalid"))
		return
	}
	if errors.Is(err, model.ErrWrongEmailOrPassword) {
		logger.Error(ctx, "[CustomerSignIn] wrong email or password", "email", req.Email)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("wrong email or password"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CustomerSignIn] failed to sign in", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to sign in"))
		return
	}

	token, err := middleware.GenerateToken(admin.Id, middleware.RoleCustomer, req.Alias, i.srv.GetJWTSecret())
	if err != nil {
		logger.Error(ctx, "[CustomerSignIn] failed to generate token", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to generate token"))
		return
	}

	logger.Info(ctx, "[CustomerSignIn] customer signed in", "alias", req.Alias, "email", req.Email)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &model.CustomerTokenResponse{
		Token: token,
	})
}
