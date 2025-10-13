package http

import (
	"failiverCheck/internal/app/dto"
	"fmt"
	"net/http"

	"failiverCheck/internal/pkg/jwtUtils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) AuthoMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtTokenStr, err := jwtUtils.ParseJWTFormHeader(ctx.Request.Header)
		if err != nil {
			h.errorHandler(ctx, 401, fmt.Errorf("access token not found"))
			return
		}
		user, err := h.UseCase.ValidateUser(ctx.Request.Context(), jwtTokenStr)
		if err != nil {
			logrus.Error(err)
			h.errorHandler(ctx, 401, err)
			return
		}
		ctx.Set("userDTO", user)
		ctx.Set("jwtToken", jwtTokenStr)
		ctx.Next()
	}
}

func (h *Handler) ModeratorValidateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtTokenStr, err := jwtUtils.ParseJWTFormHeader(ctx.Request.Header)
		if err != nil {
			h.errorHandler(ctx, 401, fmt.Errorf("access token not found"))
			return
		}
		user, err := h.UseCase.ValidateUser(ctx.Request.Context(), jwtTokenStr)
		if err != nil {
			logrus.Error(err)
			h.errorHandler(ctx, 401, err)
			return
		}
		logrus.Error(user.ID, user.IsModerator, user.Login)
		if !user.IsModerator {
			h.errorHandler(ctx, http.StatusForbidden, fmt.Errorf("not allowed role"))
		}
		ctx.Set("userDTO", user)
		ctx.Set("jwtToken", jwtTokenStr)
		ctx.Next()
	}
}

func (h *Handler) SystemCalcAccessMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := h.GetUserDTO(ctx)
		if err != nil {
			h.errorHandler(ctx, http.StatusNotFound, err)
			return
		}
		// if user.IsModerator {
		// 	ctx.Next()
		// 	return
		// }
		id := h.getIntParam(ctx, "id")
		if ctx.IsAborted() {
			return
		}
		sysCalc, err := h.UseCase.Postgres.GetSystemCalcById(uint(id))
		if err != nil {
			h.errorHandler(ctx, http.StatusNotFound, err)
			return
		}
		if sysCalc.UserID != uint(user.ID) {
			h.errorHandler(ctx, http.StatusForbidden, fmt.Errorf("access denied"))
			return
		}
		ctx.Next()
	}
}
func (h *Handler) GetUserID(ctx *gin.Context) uint {
	dto, err := h.GetUserDTO(ctx)
	if err != nil {
		h.errorHandler(ctx, 404, err)
		return 0
	}
	return uint(dto.ID)
}

func (h *Handler) GetUserDTO(ctx *gin.Context) (dto.UserDTO, error) {
	userRaw, ok := ctx.Get("userDTO")
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("user not found")
	}
	user, ok := userRaw.(dto.UserDTO)
	if !ok {
		return dto.UserDTO{}, fmt.Errorf("invalid user id type: expected int, got %T", user)
	}
	return user, nil
}

func (h *Handler) GetJWTToken(ctx *gin.Context) (string, error) {
	jwtRaw, ok := ctx.Get("jwtToken")
	if !ok {
		return "", fmt.Errorf("jwt token not found")
	}
	jwtTokenStr, ok := jwtRaw.(string)
	if !ok {
		return "", fmt.Errorf("invalid jwt token type: expected string, got %T", jwtTokenStr)
	}
	return jwtTokenStr, nil
}
