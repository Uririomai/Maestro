package file

import (
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/go-chi/render"
	"net/http"
)

// UploadImageFile godoc
// @Summary Upload image
// @Tags file
// @Accept   jpeg
// @Produce  json
// @Param image formData file true "image"
// @Success      200   {object}   model.UploadFileResponse
// @Failure      400   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /file/upload-image [post]
func (i *File) UploadImageFile(w http.ResponseWriter, r *http.Request) {
	const contentTypeHeader = "content-type"

	ctx := r.Context()

	file, info, err := r.FormFile("image")
	if err != nil {
		logger.Error(ctx, "[UploadImageFile] can't get image from form", "err", err)
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("haven't image in form"))
		return
	}

	contentType := info.Header.Get(contentTypeHeader)
	id, err := i.srv.UploadImage(ctx, file, info.Size, contentType)
	if err != nil {
		logger.Error(ctx, "[UploadImageFile] can't upload image", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("can't upload image"))
		return
	}

	logger.Info(ctx, "[UploadImageFile] image uploaded", "id", id)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.UploadFileResponse{Id: id})
}
