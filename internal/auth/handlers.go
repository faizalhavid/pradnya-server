package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/faizalhavid/pradnya-server/internal/middleware"
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

// Register godoc
//
// @Summary Register user
// @Description Create new account
// @Tags Auth
// @Accept json
// @Produce json
//
// @Param request body RegisterRequest true "Register Request"
//
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} map[string]interface{}
//
// @Router /auth/register [post]
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

// Login godoc
//
// @Summary Login user
// @Description Verify User Cred
// @Tags Auth
// @Accept json
// @Produce json
//
// @Param request body LoginRequest true "Login Request"
//
// @Success 200 {object} LoginResponse
// @Failure 403 {object} map[string]interface{}
//
// @Router /auth/login [post]
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

// Me godoc
//
// @Summary Current User
// @Tags Auth
// @Security BearerAuth
//
// @Produce json
//
// @Success 200 {object} user.UserResponse
// @Router /auth/me [get]
func (h *Handler) Me(c *gin.Context) {
	userID := c.GetString(middleware.ContextUserID)

	res, err := h.service.Me(
		userID,
	)
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

// ForgotPassword godoc
//
// @Summary Forgot Password
// @Description Send reset password email
// @Tags Auth
// @Accept json
// @Produce json
//
// @Param request body ForgotPasswordRequest true "Forgot Password Request"
//
// @Success 200 {object} ForgotPasswordResponse
// @Failure 404 {object} map[string]interface{}
//
// @Router /auth/forgot-password [post]
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

// ResetPassword godoc
//
// @Summary Reset Password
// @Description Reset user password
// @Tags Auth
// @Accept json
// @Produce json
//
// @Param request body ResetPasswordRequest true "Reset Password Request"
//
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
//
// @Router /auth/reset-password [post]
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
