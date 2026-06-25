package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		shared.AppErrorResponse(c, shared.BadRequest(err.Error()))
		return
	}
	res, err := h.service.Register(req)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			shared.AppErrorResponse(c, shared.Conflict(err.Error()))
			return
		}

		shared.AppErrorResponse(c, shared.InternalServerError("Internal server error"))
		return
	}
	shared.AppSuccessResponse(c, http.StatusAccepted, res)
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		shared.AppErrorResponse(c, shared.BadRequest(err.Error()))
		return
	}
	res, err := h.service.Login(req)
	if err != nil {
		if errors.Is(err, ErrUserNotExist) {
			shared.AppErrorResponse(c, shared.NotFound("User not found in system"))
			return
		}
		if errors.Is(err, ErrInvalidCredentials) {
			shared.AppErrorResponse(c, shared.Forbidden("Invalid Credentials"))
			return
		}
		fmt.Printf("Error %s", err)
		shared.AppErrorResponse(c, shared.InternalServerError("Internal server error"))
		return
	}
	shared.AppSuccessResponse(c, http.StatusAccepted, res)
}

func (h *Handler) Me(c *gin.Context) {
	userId := c.Param("id")
	res, err := h.service.Me(userId)
	if err != nil {
		if errors.Is(err, ErrUserNotExist) {
			shared.AppErrorResponse(c, shared.NotFound("User not found in system"))
			return
		}
		fmt.Printf("Error : %s", err)
		shared.AppErrorResponse(c, shared.InternalServerError("Internal server error"))
		return
	}
	shared.AppSuccessResponse(c, http.StatusAccepted, res)
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		shared.AppErrorResponse(c, shared.BadRequest(err.Error()))
		return
	}
	res, err := h.service.ForgotPassword(req)
	if err != nil {
		if errors.Is(err, ErrUserNotExist) {
			shared.AppErrorResponse(c, shared.NotFound("User not found in system"))
			return
		}
		fmt.Printf("Error : %s", err)
		shared.AppErrorResponse(c, shared.InternalServerError("Internal server error"))
		return
	}
	shared.AppSuccessResponse(c, http.StatusAccepted, res)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		shared.AppErrorResponse(c, shared.BadRequest(err.Error()))
		return
	}

	err := h.service.ResetPassword(req)
	if err != nil {
		if errors.Is(err, ErrInvalidToken) {
			shared.AppErrorResponse(c, shared.Unauthorized("Invalid token"))
			return
		}
		fmt.Printf("Error : %s", err)
		shared.AppErrorResponse(c, shared.InternalServerError("Internal server error"))
		return
	}
	shared.AppSuccessResponse(c, http.StatusAccepted, gin.H{"message": "Password reset successful"})
}
