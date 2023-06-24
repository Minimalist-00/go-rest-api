package repository

import "go-rest-api/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error //引数のuserにemailをもつユーザーを格納｜返り値 エラーを返すときに使う
	CreateUser(user *model.User) error                   //引数のuserをDBに保存
}
