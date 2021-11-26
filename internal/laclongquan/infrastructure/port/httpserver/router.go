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
		s.app.LikeHandler,
	)
	reportCtrl := controller.NewReportCtrl(
		s.app.ReportHandler,
	)
	// ---------------------

	// ---- ROUTES ----
	r.POST("/signup", userCtrl.SignUp)
	r.POST("/login", userCtrl.SignIn)
	r.GET("/logout", s.AuthMiddleware(), userCtrl.Signout)

	userGr := r.Group("/user")
	{
		userGr.Use(s.AuthMiddleware())
		userGr.PUT("/password", userCtrl.ChangePassword)
	}

	friendGr := r.Group("/friend")
	{
		friendGr.Use(s.AuthMiddleware())
		friendGr.POST("/request/:userid", userCtrl.NewFriendRequest)
		friendGr.PUT("/request/:userid", userCtrl.UpdateFriendRequest)
	}

	blockGr := r.Group("/block")
	{
		blockGr.Use(s.AuthMiddleware())
		blockGr.POST("", userCtrl.BlockUser)
	}

	postGr := r.Group("/post")
	{
		postGr.Use(s.AuthMiddleware())
		postGr.POST("", postCtrl.CreatePost)
		postGr.PUT("/:postid", postCtrl.EditPost)
		postGr.DELETE("/:postid", postCtrl.DeletePost)

		// like
		postGr.PUT("/:postid/like", postCtrl.LikePost)
	}

	reportGr := r.Group("/report")
	{
		reportGr.Use(s.AuthMiddleware())
		reportGr.POST("", reportCtrl.Create)
	}
	// ---------------

	return r
}
