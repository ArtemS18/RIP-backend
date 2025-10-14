package http

import (
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (h *Handler) UpdateComponentsToSystemCalc(ctx *gin.Context) {
	var update schemas.UpdateComponentToSystemCalcReq
	h.validateFields(ctx, &update)
	if ctx.IsAborted() {
		return
	}
	userID := h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	new, err := h.UseCase.UpdateComponentsToSystemCalc(dto.UpdateComponentToSystemCalcDTO{
		ReplicationCount:    update.ReplicationCount,
		ComponentID:         update.ComponentID,
		SystemCalculationID: update.SystemCalculationID,
		UserID:              userID,
	})
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
func (h *Handler) DeleteComponentsToSystemCalc(ctx *gin.Context) {
	var del schemas.ComponentToSystemCalcByIdReq
	h.validateFields(ctx, &del)
	if ctx.IsAborted() {
		return
	}
	userID := h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	err := h.UseCase.DeleteComponentsToSystemCalc(dto.ComponentToSystemCalcByIdDTO{
		ComponentID:         del.ComponentID,
		SystemCalculationID: del.SystemCalculationID,
		UserID:              userID,
	})
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, nil)
}
