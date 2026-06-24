package auth

import (
	"time"

	"github.com/faizalhavid/pradnya-server/internal/shared"
	"github.com/faizalhavid/pradnya-server/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(req RegisterRequest) (*RegisterResponse, error)
	Login(req LoginRequest) (*LoginResponse, error)
	// Me(userId string) *user.UserResponse
}

type service struct {
	repo   Repository
	jwtCfg shared.JWTConfig
}

func NewService(
	repo Repository,
	jwtCfg shared.JWTConfig,
) Service {
	return &service{
		repo:   repo,
		jwtCfg: jwtCfg,
	}
}

var _ Service = (*service)(nil)

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
func (s *service) Register(req RegisterRequest) (*RegisterResponse, error) {
	existUser, err := s.repo.FindByEmail(req.Email)
	if err == nil && existUser != nil {
		return nil, ErrEmailAlreadyExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &user.User{
		Username: req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	err = s.repo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{
		user: user.UserResponse{
			ID:       newUser.ID,
			Username: newUser.Username,
			Email:    newUser.Email,
		},
	}, nil
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
func (s *service) Login(req LoginRequest) (*LoginResponse, error) {
	authenticatedUser, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, ErrUserNotExist
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(authenticatedUser.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := s.jwtCfg.GenerateToken(
		authenticatedUser.ID,
		shared.TokenPurposeAccess,
		15*time.Minute,
	)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwtCfg.GenerateToken(
		authenticatedUser.ID,
		shared.TokenPurposeRefresh,
		30*24*time.Hour,
	)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		user: user.UserResponse{
			ID:       authenticatedUser.ID,
			Username: authenticatedUser.Username,
			Email:    authenticatedUser.Email,
		},
		credentials: CredentialsData{
			AccessToken:  *accessToken,
			RefreshToken: *refreshToken,
		},
	}, nil
}
