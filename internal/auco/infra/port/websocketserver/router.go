package websocketserver

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/zola/internal/auco/infra/port/websocketserver/controller"
)

func (s *WebsocketServer) newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// controller
	wsCtrl := controller.NewWsController(s.wm)

	// routes
	router.StaticFS("/pub", http.Dir("./public"))

	wsGroup := router.Group("/ws")
	{
		wsGroup.GET("", wsCtrl.ServeWebsocket)
	}

	return router
}
