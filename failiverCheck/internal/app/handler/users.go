package handler

import (
	"failiverCheck/internal/app/dto"
	"failiverCheck/internal/app/schemas"

	"github.com/gin-gonic/gin"
)

// Register User
// @Summary       Register User
// @Description   Register User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param  credentials body schemas.UserCredentials true "User Credentials"
// @Success      201  {object}  ds.User
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /users/register [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var credentials schemas.UserCredentials
	h.validateFields(ctx, &credentials)
	if ctx.IsAborted() {
		return
	}
	user, err := h.Postgres.RegisterUser(credentials)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	h.successHandler(ctx, 201, user)

}

// Autho User
// @Summary       Autho User
// @Description   Autho User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param  credentials body schemas.UserCredentials true "User Credentials"
// @Success      204  {object}  nil
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /users/auth [post]
func (h *Handler) AuthUser(ctx *gin.Context) {
	var credentials schemas.UserCredentials
	h.validateFields(ctx, &credentials)
	if ctx.IsAborted() {
		return
	}
	resp, err := h.UseCase.Autho(credentials)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	h.successHandler(ctx, 200, resp)

}

// Logout User
// @Summary       Logout User
// @Description   Logout User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      204  {object}  nil
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /users/logout [post]
func (h *Handler) LogoutUser(ctx *gin.Context) {
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	err := h.Postgres.LogoutUser(userId)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	h.successHandler(ctx, 204, nil)

}

// Show current User
// @Summary       Show current User
// @Description   Show current User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.UserDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /users/me [get]
func (h *Handler) GetUser(ctx *gin.Context) {
	var userId uint = h.GetUserID(ctx)
	if ctx.IsAborted() {
		return
	}
	user, err := h.Postgres.GetUserById(userId)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	userDTO := dto.ToUserDTO(user)
	h.successHandler(ctx, 200, userDTO)

}

// Update current User
// @Summary       Update current User
// @Description   Update current User
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param  update body dto.UserUpdateDTO true "User Update"
// @Success      200  {object}  dto.UserDTO
// @Failure      400  {object}  schemas.Error
// @Failure      404  {object}  schemas.Error
// @Failure      500  {object}  schemas.Error
// @Router       /users/me [put]
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
	user, err := h.Postgres.UpdateUserById(userId, update)
	if err != nil {
		h.errorHandler(ctx, 401, err)
		return
	}
	userDTO := dto.ToUserDTO(user)
	h.successHandler(ctx, 200, userDTO)

}
