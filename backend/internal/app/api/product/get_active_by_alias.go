package product

import (
	"github.com/Nikita-Kolbin/Maestro/internal/app/model"
	"github.com/Nikita-Kolbin/Maestro/internal/pkg/logger"
	"github.com/go-chi/render"
	"net/http"
)

// GetActiveProductByAlias godoc
// @Summary Get active products in website by alias
// @Tags product
// @Accept   json
// @Produce  json
// @Param alias query string true "website alias"
// @Success      200   {object}   model.ProductDTOList
// @Failure      400   {object}   model.ErrorResponse
// @Failure      500   {object}   model.ErrorResponse
// @Router /product/get-active-by-alias [get]
func (i *Product) GetActiveProductByAlias(w http.ResponseWriter, r *http.Request) {
	const aliasParamKey = "alias"

	ctx := r.Context()

	alias := r.URL.Query().Get(aliasParamKey)
	if alias == "" {
		logger.Error(ctx, "[GetActiveProductByAlias] empty alias")
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, model.NewErrResp("empty alias query param"))
		return
	}

	products, err := i.srv.GetActiveProductsByAlias(ctx, alias)
	if err != nil {
		logger.Error(ctx, "[GetActiveProductByAlias] failed to get products", "err", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, model.NewErrResp("failed to get products"))
		return
	}

	logger.Info(ctx, "[GetActiveProductByAlias] products given", "alias", alias)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, model.FromProductListToDTO(products))
}
