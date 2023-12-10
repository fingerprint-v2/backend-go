package services


type AuthService interface {
}

type authServiceImpl struct {
}

func NewAuthService() AuthService {
	return &authServiceImpl{
	}
}