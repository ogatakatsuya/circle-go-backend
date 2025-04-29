package router

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"circle/pkg/middleware"
	"circle/pkg/server/controller"
)

func NewRouter(ac controller.IAuthController, pc controller.IPostController) *echo.Echo {
	e := echo.New()

	// panicが発生した場合の処理
	e.Use(echomiddleware.Recover())

	// CORSの設定
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type,Accept,Origin,x-token"},
	}))

	zapLogger, _ := zap.NewProduction()

	// 独自のログミドルウェアを追加
	e.Use(middleware.LoggingMiddleware(*zapLogger))

	e.GET("/test", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// 認証系
	auth := e.Group("/auth")
	auth.POST("/sign-in", ac.SignIn)
	auth.POST("/sign-up", ac.SignUp)
	auth.POST("/sign-out", ac.SignOut)

	// Post系
	post := e.Group("/posts")
	post.GET("", pc.GetAll)

	return e
}
