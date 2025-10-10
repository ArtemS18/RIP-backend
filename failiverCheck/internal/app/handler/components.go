package handler

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
	component, err := h.Repository.GetComponentById(id)
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
	if searchQuery == "" {
		components, err = h.Repository.GetComponents()
		if err != nil {
			h.errorHandler(ctx, http.StatusBadRequest, err)
			return
		}
	} else {
		components, err = h.Repository.GetComponentsByTitle(searchQuery)
		if err != nil {
			h.errorHandler(ctx, http.StatusNotFound, err)
			return
		}
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
	component, err := h.Repository.UpdateComponentById(uint(id), update)
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
	component, err := h.Repository.CreateComponent(create)
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
	component, err := h.Repository.GetComponentById(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	imgUrl := component.Img

	if err := h.Repository.DeletedComponentById(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	if err := h.Repository.DeleteComponentImg(ctx, &imgUrl); err != nil {
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
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	err = h.Repository.AddComponentInSystemCalc(uint(id), userId)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *Handler) UpdateComponentImg(ctx *gin.Context) {
	id := h.getIntParam(ctx, "id")
	if ctx.IsAborted() {
		return
	}
	contentType, fileSize := h.getFileHeaders(ctx)
	if ctx.IsAborted() {
		return
	}
	location, err := h.Repository.UploadComponentImg(ctx, dto.ComponentImgCreateDTO{
		File:        ctx.Request.Body,
		FilePath:    "img/",
		FileSize:    fileSize,
		ContentType: contentType,
	})
	log.Info(location)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}
	component, err := h.Repository.GetComponentById(id)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return
	}
	if component.Img != "" {
		if err = h.Repository.DeleteComponentImg(ctx, &component.Img); err != nil {
			log.Error(err)
		}
	}
	h.Repository.UpdateComponentById(uint(id), dto.UpdateComponentDTO{Img: &location})
	h.successHandler(ctx, 200, map[string]string{"img_url": location})
}
