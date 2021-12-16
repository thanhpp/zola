package httpserver

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/port/httpserver/controller"
)

func (s HTTPServer) formURL() string {
	return "http://" + s.cfg.Host + ":" + s.cfg.Port
}

func (s HTTPServer) formMediaURL(post entity.Post, media entity.Media) string {
	return fmt.Sprintf("%s/post/%s/media/%s", s.formURL(), post.ID(), media.ID())
}

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
		s.formMediaURL,
	)
	reportCtrl := controller.NewReportCtrl(
		s.app.ReportHandler,
	)
	// ---------------------

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
		postGr.GET("/:postid", postCtrl.GetPost)
		postGr.POST("", postCtrl.CreatePost)
		postGr.PUT("/:postid", postCtrl.EditPost)
		postGr.DELETE("/:postid", postCtrl.DeletePost)

		// like
		postGr.PUT("/:postid/like", postCtrl.LikePost)

		// comment
		postGr.POST("/:postid/comment", postCtrl.CreateComment)
		postGr.PUT("/:postid/comment/:commentid", postCtrl.UpdateComment)
		postGr.DELETE("/:postid/comment/:commentid", postCtrl.DeleteComment)

		// media
		postGr.GET("/:postid/media/:mediaid", postCtrl.GetMedia)
	}

	reportGr := r.Group("/report")
	{
		reportGr.Use(s.AuthMiddleware())
		reportGr.POST("", reportCtrl.Create)
	}
	// ---------------

	return r
}
