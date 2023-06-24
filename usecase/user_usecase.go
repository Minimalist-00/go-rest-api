package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error) //JWTを返すためにstring型
}

type userUsecase struct {
	userRepo repository.IUserRepository
}

// func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
// 	return &userUsecase{
// 		userRepo: ur,
// 	}
// }
