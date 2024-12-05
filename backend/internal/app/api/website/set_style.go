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

// SetStyle godoc
// @Summary Set new website styles
// @Security ApiKeyAuth
// @Tags website
// @Accept   json
// @Produce  json
// @Param input body model.SetWebsiteStyleRequest true "styles and alias"
// @Success      200   {object}   model.WebsiteStylesDTO
// @Failure      400   {object}   model.ErrorResponse
// @Failure      403   {object}   model.ErrorResponse
// @Failure      404   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /website/set-style [post]
func (i *Website) SetStyle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsAdmin(ctx) {
		logger.Error(ctx, "[SetStyle] user is not admin")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	adminId := middleware.GetUserId(ctx)

	var req model.SetWebsiteStyleRequest

	err := render.DecodeJSON(r.Body, &req)
	if errors.Is(err, io.EOF) {
		logger.Error(ctx, "[SetStyle] request body is empty")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty request"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[SetStyle] failed to decode request body", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("failed to decode request"))
		return
	}

	website, err := i.srv.GetWebsiteByAlias(ctx, req.WebsiteAlias)
	if errors.Is(err, model.ErrNotFound) {
		logger.Error(ctx, "[SetStyle] website not found", "alias", req.WebsiteAlias)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, model.NewErrResp("website not found"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[SetStyle] failed to check admin", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to check admin"))
		return
	}
	if website.AdminId != adminId {
		logger.Error(ctx, "[SetStyle] admin is not owner")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	sections, err := i.srv.SetWebsiteStyle(ctx, req.WebsiteAlias, model.FromSetWebsiteStyleRequestToSections(&req))
	if err != nil {
		logger.Error(ctx, "[SetStyle] failed to set styles", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to set styles"))
		return
	}

	logger.Info(ctx, "[SetStyle] style is set", "alias", req.WebsiteAlias)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromSectionsToDTO(sections))
}
