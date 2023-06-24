package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error) //JWTを返すためにstring型
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	//パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10) //GenerateFromPassword関数により、パスワードをハッシュ化
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)} //ハッシュ化したパスワードをnewUserに格納
	if err := uu.ur.CreateUser(&newUser); err != nil {               //引数のnewUserをDBに保存
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	} //レスポンス用のUserResponse型の変数を作成
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	return "", nil
}
