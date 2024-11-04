package website

import (
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// CreateWebsite godoc
// @Summary Create website
// @Security ApiKeyAuth
// @Tags website
// @Accept json
// @Produce  json
// @Param input body model.CreateWebsiteRequest true "alias new website"
// @Success      200   {object}   nil
// @Failure      400   {object}   model.ErrorResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      409   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /website/create [post]
func (i *Website) CreateWebsite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsAdmin(ctx) {
		logger.Error(ctx, "[CreateWebsite] user is not admin")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	adminId := middleware.GetUserId(ctx)

	var req model.CreateWebsiteRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[CreateWebsite] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CreateWebsite] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	err = i.srv.CreateWebsite(ctx, req.Alias, adminId)
	if errors.Is(err, model.ErrAdminHaveWebsite) {
		logger.Error(ctx, "[CreateWebsite] admin", "admin already have website", "id", adminId)
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, model.NewErrResp("admin already have website"))
		return
	}
	if errors.Is(err, model.ErrAliasTaken) {
		logger.Error(ctx, "[CreateWebsite] alias is already taken", "alias", req.Alias)
		render.Status(r, http.StatusConflict)
		render.JSON(w, r, model.NewErrResp("alias is already taken"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[CreateWebsite] failed to create website", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to create website"))
		return
	}

	logger.Info(ctx, "[CreateWebsite] website created", "alias", req.Alias)

	// TODO: возвращать id и alias
	render.Status(r, http.StatusOK)
	render.JSON(w, r, nil)
}
