package shared

type JWTClaims interface {
	GenerateJWT() (string, error)
	ValidateJWT(token string) (bool, error)
}
