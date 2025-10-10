package handler

import (
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterUser(ctx *gin.Context) {
	var credentials schemas.UserCredentials
	h.validateFields(ctx, &credentials)
	if ctx.IsAborted() {
		return
	}
	user, err := h.Repository.RegisterUser(credentials)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	h.successHandler(ctx, 201, user)

}

func (h *Handler) AuthUser(ctx *gin.Context) {
	var credentials schemas.UserCredentials
	h.validateFields(ctx, &credentials)
	if ctx.IsAborted() {
		return
	}
	err := h.Repository.AuthUser(credentials)
	if err != nil {
		h.errorHandler(ctx, 401, err)
	}
	h.successHandler(ctx, 204, nil)

}

func (h *Handler) LogoutUser(ctx *gin.Context) {
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	err := h.Repository.LogoutUser(userId)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	h.successHandler(ctx, 204, nil)

}

func (h *Handler) GetUser(ctx *gin.Context) {
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	user, err := h.Repository.GetUserById(userId)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	userDTO := dto.ToUserDTO(user)
	h.successHandler(ctx, 200, userDTO)

}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	var update dto.UserUpdateDTO
	h.validateFields(ctx, &update)
	if ctx.IsAborted() {
		return
	}
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	user, err := h.Repository.UpdateUserById(userId, update)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	userDTO := dto.ToUserDTO(user)
	h.successHandler(ctx, 200, userDTO)

}
