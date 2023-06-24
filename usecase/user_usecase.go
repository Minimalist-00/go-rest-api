package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	resUser := model.UserResponse{ //レスポンス用のUserResponse型の変数を作成
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	//クライアントからのEmailがDB内に存在するかを確認
	storedUser := model.User{} //DBから取得したユーザー情報を格納するための変数
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// ハッシュ化されたパスと元のパスの一致を比較
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //JWTの生成
		"user_id": storedUser.ID,                         //ユーザーIDの設定
		"exp":     time.Now().Add(time.Hour * 12).Unix(), //有効期限の設定
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) //著名済みの文字列を生成 <- トークンとして使うことで、信頼性↑
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
