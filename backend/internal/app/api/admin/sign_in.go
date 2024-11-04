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

// AdminSignIn godoc
// @Summary      SingIn admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        input body       model.AdminEmailPasswordRequest true "sign in"
// @Success      200   {object}   model.ErrorResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router       /admin/sign-in [post]
func (i *Admin) AdminSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.AdminEmailPasswordRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[AdminSignIn] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[AdminSignIn] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	admin, err := i.srv.GetAdminIdByEmailPassword(ctx, req.Email, req.Password)
	if errors.Is(err, model.ErrInvalidEmail) {
		logger.Error(ctx, "[AdminSignIn] email is invalid", "email", req.Email)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("email is invalid"))
		return
	}
	if errors.Is(err, model.ErrWrongEmailOrPassword) {
		logger.Error(ctx, "[AdminSignIn] wrong email or password", "email", req.Email)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("wrong email or password"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[AdminSignIn] failed to sign in", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to sign in"))
		return
	}

	token, err := middleware.GenerateToken(admin.Id, middleware.RoleAdmin, "", i.srv.GetJWTSecret())
	if err != nil {
		logger.Error(ctx, "[AdminSignIn] failed to generate token", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to generate token"))
		return
	}

	logger.Info(ctx, "[AdminSignIn] admin signed in", "email", req.Email)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &model.AdminTokenResponse{
		Token: token,
	})
}
