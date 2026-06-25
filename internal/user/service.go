package user

type Service interface {
	GetUserByID(id string) (*UserResponse, error)
	UpdateProfile(id string, profile *ProfileRequest) (*UserResponse, error)
	DeActivateAccount(id string) error
	GetProfile(id string) (*ProfileResponse, error)
	GetImageProfile(id string) (string, error)
}

type service struct {
	repo Repository
}
