package auth

import (
	"net/http"

	"github.com/faizalhavid/pradnya-server/internal/shared"
)

var ErrEmailAlreadyExists = shared.NewAppError(
	http.StatusConflict,
	"EMAIL_ALREADY_EXISTS",
	"email already exists",
)

var ErrInvalidCredentials = shared.NewAppError(
	http.StatusUnauthorized,
	"INVALID_CREDENTIALS",
	"invalid email or password",
)

var ErrInvalidToken = shared.NewAppError(
	http.StatusUnauthorized,
	"INVALID_TOKEN",
	"invalid token",
)

var ErrUserNotExist = shared.NewAppError(http.StatusNotFound, "USER_NOT_FOUND", "User Not Found")
