package http

import (
	"failiverCheck/internal/app/ds"
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"
	"net/http"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

// Show a component by id
// @Summary       Show a component by id
// @Description   Show a component by id
// @Tags         Components
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "Component ID"
// @Success      200  {object}  ds.Component
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /components/{id} [get]
func (h *Handler) GetComponent(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	log.Info(id)
	component, err := h.UseCase.GetComponent(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, component)
}

// Show a components list
// @Summary       Show a components list
// @Description   Show a components list
// @Tags         Components
// @Accept       json
// @Produce      json
// @Success      200  {array}  ds.Component
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /components/ [get]
func (h *Handler) GetComponents(ctx *gin.Context) {
	var components []ds.Component
	var err error
	searchQuery := ctx.Query("search")

	log.Info(searchQuery)
	components, err = h.UseCase.GetComponents(searchQuery)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	h.successHandler(ctx, http.StatusOK, schemas.ComponentsRes{Components: components})
}

// Update a component by id
// @Summary       Show a component by id
// @Description   Show a component by id
// @Tags         Components
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "Component ID"
// @Param  update body dto.UpdateComponentDTO true "Update component schema"
// @Success      200  {object}  ds.Component
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /components/{id} [get]
func (h *Handler) UpdateComponent(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	var update dto.UpdateComponentDTO
	if err := ctx.BindJSON(&update); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	component, err := h.UseCase.UpdateComponent(uint(id), update)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, component)
}

// Create a component
// @Summary       Create a component
// @Description   Create a component
// @Tags         Components
// @Accept       json
// @Produce      json
// @Param  create body dto.CreateComponentDTO true "Create component schema"
// @Security     BearerAuth
// @Success      200  {object}  ds.Component
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /components/ [post]
func (h *Handler) CreateComponent(ctx *gin.Context) {
	var create dto.CreateComponentDTO
	h.validateFields(ctx, &create)
	if ctx.IsAborted() {
		return
	}
	component, err := h.UseCase.CreateComponent(create)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, http.StatusOK, component)
}

// Delete a component by id
// @Summary       Delete a component by id
// @Description   Delete a component by id
// @Tags         Components
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "Component ID"
// @Security     BearerAuth
// @Success      204  {object}  nil
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /components/{id} [delete]
func (h *Handler) DeleteComponent(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	err := h.UseCase.DeleteComponent(uint(id))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// Add component in system caclulation
// @Summary       Add component in system caclulation
// @Description   Add component in system caclulation
// @Tags         Components
// @Accept       json
// @Produce      json
// @Param        id    path     int  true "Component ID"
// @Security     BearerAuth
// @Success      204  {object}  nil
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /components/{id}/system_calc/ [post]
func (h *Handler) AddComponentInSystemCalc(ctx *gin.Context) {
	var err error
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	log.Info(userId)
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	err = h.UseCase.AddComponentInSystemCalc(userId, uint(id))
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// Update component image
// @Summary       Update component image
// @Description   Update component image
// @Tags         Components
// @Accept       multipart/form-data
// @Produce      json
// @Param        id          path     int                         true "Component ID"
// @Security     BearerAuth
// @Param         img formData file true "image"
// @Success      200         {object}  map[string]string
// @Failure      400         {object}  schemas.Error
// @Failure      404         {object}  schemas.Error
// @Failure      500         {object}  schemas.Error
// @Router       /components/{id}/img [post]
func (h *Handler) UpdateComponentImg(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	contentType, fileSize := h.getFileHeaders(ctx)
	if ctx.IsAborted() {
		return
	}
	location, err := h.UseCase.UpdateComponentImg(uint(id), dto.ComponentImgCreateDTO{
		File:        ctx.Request.Body,
		FilePath:    "img/",
		FileSize:    fileSize,
		ContentType: contentType,
	})
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	h.successHandler(ctx, 200, map[string]string{"img_url": location})
}
