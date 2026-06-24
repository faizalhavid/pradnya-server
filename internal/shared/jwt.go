package shared

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPurpose string

const (
	TokenPurposeAccess  TokenPurpose = "access"
	TokenPurposeRefresh TokenPurpose = "refresh"
	TokenPurposeReset   TokenPurpose = "reset_password"
	TokenPurposeVerify  TokenPurpose = "verify_email"
)

type JWTConfig struct {
	Secret string
	Issuer string
}

type TokenData struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Claims struct {
	UserID string       `json:"user_id"`
	Type   TokenPurpose `json:"type"`

	jwt.RegisteredClaims
}

func (c JWTConfig) GenerateToken(
	userId string,
	purpose TokenPurpose,
	duration time.Duration,
) (*TokenData, error) {
	now := time.Now()
	ExpiredAt := now.Add(duration)

	claims := Claims{
		UserID: userId,
		Type:   purpose,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    c.Issuer,
			ExpiresAt: jwt.NewNumericDate(ExpiredAt),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(
		[]byte(c.Secret),
	)
	if err != nil {
		return nil, err
	}
	return &TokenData{
		Token:     tokenString,
		ExpiredAt: ExpiredAt,
	}, nil
}
