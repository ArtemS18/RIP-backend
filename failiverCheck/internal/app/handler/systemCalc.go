package handler

import (
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Show an system calc
// @Summary      Show an system calc
// @Description  get system cacl by id
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "System Cacl ID"
// @Success      200  {object}  dto.SystemCalculationDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/{id} [get]
func (h *Handler) GetSystemCalc(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	systemCalc, err := h.Repository.GetSystemCalcById(uint(id))
	sysCalcsResp := dto.ToSystemCalculationDTO(systemCalc)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}

	h.successHandler(ctx, 200, sysCalcsResp)
}

// Show an system calc list
// @Summary      Show an system calc list
// @Description  get system cacl in
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Param        state    query     string  false  "name search by q"
// @Param        date_formed_start    query     string  false  "2006-02-01"
// @Param        date_formed_end    query     string  false  "2006-02-01"
// @Success      200  {array}   dto.SystemCalculationInfoDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/ [get]
func (h *Handler) GetSystemCalcList(ctx *gin.Context) {
	var filters dto.SystemCalcFilters
	if errQ := ctx.BindQuery(&filters); errQ != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errQ)
		return
	}
	sysCalcs, err := h.Repository.GetSystemCalcList(filters)
	sysCalcsResp := dto.ToSystemCalculationInfoListDTO(sysCalcs)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	h.successHandler(ctx, http.StatusOK, sysCalcsResp)
}

// Show user bucket
// @Summary      Show user bucket
// @Description  get user bucket system cac
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Success      200  {object}   dto.CurrentUserBucketDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/my_bucket [get]
func (h *Handler) GetSystemCalcBucket(ctx *gin.Context) {
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	bucket, err := h.Repository.GetCurrentSysCalcAndCount(userId)
	if err != nil {
		h.errorHandler(ctx, 400, err)
		return
	}
	h.successHandler(ctx, http.StatusOK, bucket)
}

// Update system calc by id
// @Summary      Update system calc by id
// @Description  update system cacl by id
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "System Cacl ID"
// @Param        update body schemas.UpdateSystemCalcFields true "Update schema"
// @Success      200  {object}   dto.SystemCalculationDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/{id} [put]
func (h *Handler) UpdateSystemCalc(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var data schemas.UpdateSystemCalcFields
	h.validateFields(ctx, &data)
	if ctx.IsAborted() {
		return
	}
	system, err := h.Repository.UpdateSystemCalc(uint(sysCalcId), dto.UpdateSystemCalcDTO{SystemName: data.SystemName})
	systemDto := dto.ToSystemCalculationDTO(system)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 200, systemDto)
}

// Update system calc status to FORMED by id
// @Summary       Update system calc status to FORMED by id
// @Description   Update system calc status to FORMED by id
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "System Cacl ID"
// @Success      200  {object}   dto.SystemCalculationDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/{id}/status_formed [put]
func (h *Handler) UpdateSystemCalcStatusToFormed(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	system, err := h.Repository.UpdateSystemCalcStatusToFormed(uint(sysCalcId))
	dto := dto.ToSystemCalculationDTO(system)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 200, dto)
}

// Moderate system calc, update status by id
// @Summary       Moderate system calc
// @Description   Moderate system calc, update status by id
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "System Cacl ID"
// @Success      200  {object}   dto.SystemCalculationDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/{id}/status [put]
func (h *Handler) UpdateSystemCalcStatusModerator(ctx *gin.Context) {
	moderatorId := uint(1)
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var data schemas.UpdateSystemCalcStatus
	h.validateFields(ctx, &data)
	if ctx.IsAborted() {
		return
	}
	system, err := h.Repository.UpdateSystemCalcStatusModerator(uint(sysCalcId), moderatorId, data.Command)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	dto := dto.ToSystemCalculationDTO(system)
	h.successHandler(ctx, 200, dto)
}

// Delete system calc by id
// @Summary       Delete system calc by id
// @Description   Delete system calc by id
// @Tags         System caclculation
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "System Cacl ID"
// @Success      204  {object}  schemas.OKResponse
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /system_calcs/{id} [delete]
func (h *Handler) DeleteSystemCalc(ctx *gin.Context) {
	sysCalcId := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	err := h.Repository.DeleteSystemCalc(uint(sysCalcId))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, 204, nil)
}
