package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(status int, code string, mesage string) *AppError {
	return &AppError{Status: status, Code: code, Message: mesage}
}

func AppSuccessResponse(
	c *gin.Context,
	status int,
	data any,
) {
	c.JSON(status, gin.H{
		"success": true,
		"data":    data,
	})
}

func AppErrorResponse(
	c *gin.Context,
	err error,
) {
	appErr, ok := err.(*AppError)
	if ok {
		c.JSON(appErr.Status, gin.H{
			"succss":  false,
			"code":    appErr.Code,
			"message": appErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"code":    "INTERNAL_SERVER_ERROR",
		"message": "internal server error",
	})
}

// functions for common app error
func BadRequest(message string) *AppError {
	return NewAppError(http.StatusBadRequest, "BAD_REQUEST", message)
}

func Unauthorized(message string) *AppError {
	return NewAppError(http.StatusUnauthorized, "UNAUTHORIZED", message)
}

func Forbidden(message string) *AppError {
	return NewAppError(http.StatusForbidden, "FORBIDDEN", message)
}

func NotFound(message string) *AppError {
	return NewAppError(http.StatusNotFound, "NOT_FOUND", message)
}

func Conflict(message string) *AppError {
	return NewAppError(http.StatusConflict, "CONFLICT", message)
}

func InternalServerError(message string) *AppError {
	return NewAppError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message)
}
