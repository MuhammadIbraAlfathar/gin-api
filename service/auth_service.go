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
