package website

import (
	"errors"
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/middleware"
	"github.com/go-chi/render"
	"net/http"
)

// GetMyWebsite godoc
// @Summary Get admins website
// @Security ApiKeyAuth
// @Tags website
// @Accept   json
// @Produce  json
// @Success      200   {object}   model.WebsiteDTO
// @Failure      403   {object}   model.ErrorResponse
// @Failure      404   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /website/get-my-website [get]
func (i *Website) GetMyWebsite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !middleware.IsAdmin(ctx) {
		logger.Error(ctx, "[GetMyWebsite] user is not admin")
		render.Status(r, http.StatusForbidden)
		render.JSON(w, r, model.NewErrResp("permission denied"))
		return
	}

	adminId := middleware.GetUserId(ctx)

	website, err := i.srv.GetWebsiteByAdminId(ctx, adminId)
	if errors.Is(err, model.ErrWebsiteNotFound) {
		logger.Error(ctx, "[GetMyWebsite] admin", "website not found", "id", adminId)
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, model.NewErrResp("website not found"))
		return
	}
	if err != nil {
		logger.Error(ctx, "[GetMyWebsite] failed to create website", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to create website"))
		return
	}

	logger.Info(ctx, "[GetMyWebsite] website given", "alias", website.Alias)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromWebsiteToDTO(website))
}
