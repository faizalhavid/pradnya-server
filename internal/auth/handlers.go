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

		shared.AppErrorResponse(c, shared.InternalServerError("internal server error"))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "register success", "data": res})
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
		shared.AppErrorResponse(c, shared.InternalServerError("internal server error"))
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "login success", "data": res})
}
