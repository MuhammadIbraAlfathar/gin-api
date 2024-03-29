package service

import (
	"github.com/MuhammadIbraAlfathar/gin-api/dto"
	"github.com/MuhammadIbraAlfathar/gin-api/entity"
	"github.com/MuhammadIbraAlfathar/gin-api/errorhandler"
	"github.com/MuhammadIbraAlfathar/gin-api/helper"
	"github.com/MuhammadIbraAlfathar/gin-api/repository"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {

	// validasi email jika sudah terdaftar
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{
			Message: "Email already registered",
		}
	}

	// validasi password dengan password_confirmation
	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{
			Message: "Password not match",
		}
	}

	hashPassword, err := helper.HashPassword(req.Password)

	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
		Gender:   req.Gender,
	}

	if err = s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil

}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotfoundError{
			Message: "wrong email or password",
		}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotfoundError{
			Message: "wrong email or password",
		}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil

}
