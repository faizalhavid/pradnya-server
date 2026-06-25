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
	Me(userId string) (*user.UserResponse, error)
	ForgotPassword(req ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ResetPassword(Req ResetPasswordRequest) error
}

type service struct {
	repo   Repository
	jwtCfg shared.JWTConfig
	mailer shared.Mailer
}

func NewService(
	repo Repository,
	jwtCfg shared.JWTConfig,
	mailer shared.Mailer,
) Service {
	return &service{
		repo:   repo,
		jwtCfg: jwtCfg,
		mailer: mailer,
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

// Me godoc
//
// @Summary Me user
// @Description User
// @Tags Auth
// @Accept json
// @Produce json
//
// @Param id path string true "User ID"
//
// @Success 200 {object} user.UserResponse
// @Failure 403 {object} map[string]interface{}
//
// @Router /auth/me [post]
func (s *service) Me(userId string) (*user.UserResponse, error) {
	foundedUser, err := s.repo.FindByID(userId)
	if err != nil {
		return nil, ErrUserNotExist
	}
	return &user.UserResponse{
		ID:       foundedUser.ID,
		Username: foundedUser.Username,
		Email:    foundedUser.Email,
	}, nil
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
func (s *service) ForgotPassword(req ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	foundedUser, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, ErrUserNotExist
	}
	resetToken, err := s.jwtCfg.GenerateToken(
		foundedUser.ID,
		shared.TokenPurposeReset,
		15*time.Minute,
	)
	if err != nil {
		return nil, err
	}
	resetLink := "https://your-frontend-app.com/reset-password?token=" + resetToken.Token
	subject := "Reset Password"
	body := "Hello " + foundedUser.Username + ",\n\n" +
		"We received a request to reset your password. Please click the link below to reset your password:\n\n" +
		resetLink + "\n\n" +
		"If you did not request a password reset, please ignore this email.\n\n" +
		"Best regards,\n" +
		"Your App Team"

	err = s.mailer.SendMail(foundedUser.Email, subject, body)
	if err != nil {
		return nil, err
	}

	return &ForgotPasswordResponse{
		ToEmail:    foundedUser.Email,
		ResetToken: *resetToken,
	}, nil
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
func (s *service) ResetPassword(req ResetPasswordRequest) error {
	claims, err := s.jwtCfg.ValidateToken(req.Token)
	if err != nil {
		return err
	}
	if claims.Type != shared.TokenPurposeReset {
		return ErrInvalidToken
	}

	foundedUser, err := s.repo.FindByID(claims.UserID)
	if err != nil {
		return ErrUserNotExist
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	foundedUser.Password = string(hashedPassword)
	err = s.repo.UpdatePassword(foundedUser.ID, foundedUser.Password)
	if err != nil {
		return err
	}

	return nil
}
