package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/controller"
)

func (s *HTTPServer) newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// ---- CONTROLLERS ----
	userCtrl := controller.NewUserCtrl(
		s.app.UserHandler,
		*s.auth,
	)
	// ---------------------

	// ---- ROUTES ----
	r.POST("/signup", userCtrl.SignUp)
	r.POST("/signin", userCtrl.SignIn)
	r.GET("/signout", s.AuthMiddleware(), userCtrl.Signout)
	// ---------------

	return r
}
