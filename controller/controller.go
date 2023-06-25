package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type UserController struct {
	uu usecase.IUserUsecase
}

// usecaseを「Dependency Injection」するための関数（コンストラクタ）
// usecaseのインスタンスを受け取る
func NewUserController(uu usecase.IUserUsecase) IController {
	return &UserController{uu}
}

func (uc *UserController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil { //リクエストボディをuserにバインド（User型に変換して格納）
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user) //usecaseのSignUpメソッドを呼び出し
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}
