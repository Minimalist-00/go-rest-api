package router

import (
	"go-rest-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	// ログイン関係のエンドポイントにの設定
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	t := e.Group("/tasks") // タスク関係のエンドポイントのグループ化
	// ミドルウェアの設定
	t.Use(echojwt.WithConfig(echojwt.Config{ //エンドポイントにミドルウェアの追加
		SigningKey:  []byte(os.Getenv("SECRET")), // 環境変数からシークレットキーを取得
		TokenLookup: "cookie:token",              // cookieからトークンを取得
	}))
	// タスク関係のエンドポイントの設定
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
