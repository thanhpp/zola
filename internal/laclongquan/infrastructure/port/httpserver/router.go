package httpserver

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
		s.app.PostHandler,
		*s.auth,
		s.resolveMediaURL,
		s.formUserMediaURL,
		s.resolveUserMediaURL,
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
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	})
	r.POST("/signup", userCtrl.SignUp)
	r.POST("/login", userCtrl.SignIn)
	r.GET("/logout", s.AuthMiddleware(), userCtrl.Signout)

	userGr := r.Group("/user")
	{
		userGr.Use(s.AuthMiddleware())
		userGr.GET("/:userid", userCtrl.GetUserInfo)
		userGr.GET("/:userid/media/:mediaid", userCtrl.GetUserMedia)

		userGr.PUT("", userCtrl.SetUserInfo)
		userGr.PUT("/password", userCtrl.ChangePassword)
	}

	friendGr := r.Group("/friend")
	{
		friendGr.Use(s.AuthMiddleware())
		friendGr.GET("", userCtrl.GetFriends)
		friendGr.GET("/:userid", userCtrl.GetFriends)
		friendGr.GET("/requested", userCtrl.GetRequestedFriends)
		friendGr.GET("/requested/:userid", userCtrl.GetRequestedFriends)

		friendGr.POST("/request/:userid", userCtrl.NewFriendRequest)
		friendGr.PUT("/request/:userid", userCtrl.UpdateFriendRequest)
	}

	blockGr := r.Group("/block")
	{
		blockGr.Use(s.AuthMiddleware())
		blockGr.POST("", userCtrl.BlockUser)

		blockGr.POST("/diary", userCtrl.BlockDiary)
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

	adminGr := r.Group("/admin")
	{
		adminGr.Use(s.AuthMiddleware())
		adminGr.GET("/users", userCtrl.AdminGetUsers)
	}
	// ---------------

	return r
}
