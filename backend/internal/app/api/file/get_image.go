package file

import (
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

// GetImageFile godoc
// @Summary Get image by id
// @Tags     file
// @Accept   json
// @Produce  json
// @Param    image-id  path  string  true  "image id"
// @Success      200   {object}   []byte
// @Failure      403   {object}   model.ErrorResponse
// @Failure      404   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /file/get-image/{image-id} [get]
func (i *File) GetImageFile(w http.ResponseWriter, r *http.Request) {
	const (
		contentTypeHeader = "content-type"
		imageIdURLParam   = "image-id"
	)

	ctx := r.Context()

	imageId := chi.URLParam(r, imageIdURLParam)
	reader, contentType, err := i.srv.GetImage(ctx, imageId)
	if err != nil {
		logger.Error(ctx, "[GetImageFile] failed to get image", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get image"))
		return
	}

	buffer, err := io.ReadAll(reader)
	if err != nil {
		logger.Error(ctx, "[GetImageFile] failed create buffer", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get image"))
		return
	}

	_, err = w.Write(buffer)
	if err != nil {
		logger.Error(ctx, "[GetImageFile] failed to write response", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get image"))
		return
	}

	w.Header().Set(contentTypeHeader, contentType)

	logger.Info(ctx, "[GetImageFile] image given", "id", imageId)

	render.Status(r, http.StatusOK)
}
