package admin

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/render"
	
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
)

// SignUp godoc
// @Summary      SingUp admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        input body       model.AdminEmailPasswordRequest true "sign up"
// @Success      200   {object}   model.ErrorResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      409   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router       /admin/sign-up [post]
func (i *Admin) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.AdminEmailPasswordRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[SignUp] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[SignUp] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	id, err := i.srv.CreateAdmin(ctx, req.Email, req.Password)
	if errors.Is(err, model.ErrInvalidEmail) {
		logger.Error(ctx, "[SignUp] email is invalid", "email", req.Email)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("email is invalid"))
		return
	}
	if errors.Is(err, model.ErrEmailRegistered) {
		logger.Error(ctx, "[SignUp] email is already taken", "email", req.Email)
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, model.NewErrResp("email is already taken"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[SignUp] failed to create admin", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to create admin"))
		return
	}

	token, err := middleware.GenerateToken(id, middleware.RoleAdmin, "", i.srv.GetJWTSecret())
	if err != nil {
		logger.Error(ctx, "[SignUp] failed to generate token", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to generate token"))
		return
	}

	logger.Info(ctx, "[SignUp] admin created", "email", req.Email)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &model.AdminTokenResponse{
		Token: token,
	})
}
