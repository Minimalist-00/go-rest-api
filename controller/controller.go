package controller

import (
	"go-rest-api/usecase"

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
