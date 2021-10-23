package httpserver

import "github.com/gin-gonic/gin"

func newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.POST("/signup")

	return r
}
