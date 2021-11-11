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
	postCtrl := controller.NewPostCtrl(
		s.app.PostHandler,
	)
	reportCtrl := controller.NewReportCtrl(
		s.app.ReportHandler,
	)
	// ---------------------

	// ---- ROUTES ----
	r.POST("/signup", userCtrl.SignUp)
	r.POST("/login", userCtrl.SignIn)
	r.GET("/logout", s.AuthMiddleware(), userCtrl.Signout)

	postGr := r.Group("/post")
	{
		postGr.Use(s.AuthMiddleware())
		postGr.POST("", postCtrl.CreatePost)
		postGr.PUT("/:postid", postCtrl.EditPost)
	}

	reportGr := r.Group("/report")
	{
		reportGr.Use(s.AuthMiddleware())
		reportGr.POST("", reportCtrl.Create)
	}
	// ---------------

	return r
}
