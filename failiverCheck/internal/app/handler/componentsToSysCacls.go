package handler

import (
	"failiverCheck/internal/app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Update component in system calc
// @Summary       Update component in system calc
// @Description   Update component in system calc
// @Tags         Components To SystemCalc
// @Accept       json
// @Produce      json
// @Param  create body dto.UpdateComponentToSystemCalcDTO true "Update schema"
// @Success      200  {object}  dto.ComponentsToSystemCalcDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs_to_components/ [put]
func (h *Handler) UpdateComponentsToSystemCac(ctx *gin.Context) {
	var update dto.UpdateComponentToSystemCalcDTO
	h.validateFields(ctx, &update)
	if ctx.IsAborted() {
		return
	}
	orm, err := h.Postgres.UpdateComponentsToSystemCalc(update)
	new := dto.ToComponentsToSystemCalcDTO(orm)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, new)
}

// Delete component in system calc
// @Summary       Delete component in system calc
// @Description   Delete component in system calc
// @Tags         Components To SystemCalc
// @Accept       json
// @Produce      json
// @Param  create body dto.ComponentToSystemCalcByIdDTO true "Delete schema"
// @Success      200  {object}  schemas.OKResponse
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs_to_components/ [delete]
func (h *Handler) DeleteComponentsToSystemCac(ctx *gin.Context) {
	var ids dto.ComponentToSystemCalcByIdDTO
	if err := ctx.BindJSON(&ids); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(ids); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	err := h.Postgres.DeleteComponentsToSystemCalc(ids)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, nil)
}
