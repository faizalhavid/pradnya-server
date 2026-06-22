package auth

import "github.com/faizalhavid/pradnya-server/internal/user"

type Service interface {
	Register(req RegisterRequest) (*RegisterResponse, error)
	// Login(req LoginRequest) (*LoginResponse, error)
	// Me(userId string) *user.UserResponse
}

type service struct {
	repo Repository
}

func NewService(
	repo Repository,
) Service {
	return &service{
		repo: repo,
	}
}

var _ Service = (*service)(nil)

func (s *service) Register(req RegisterRequest) (*RegisterResponse, error) {
}
