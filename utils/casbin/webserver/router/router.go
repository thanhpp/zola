package router

import (
	"time"

	"bitbucket.org/tysud/gt-casbin/webserver/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	//CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handlers := new(controller.CasbinHandler)
	r.POST("/author/check", handlers.CheckAuthorization)
	// user
	r.POST("/user/per", handlers.AddPermissionToUserOnDomain)
	r.DELETE("/user/per", handlers.DeletePermissionOfUserOnDomain)
	r.DELETE("/user", handlers.DeleteUser)

	r.POST("/user/role", handlers.AddRoleToUser)
	r.DELETE("/user/role", handlers.DeleteRoleOfUser)

	// role
	r.POST("/role/per", handlers.AddPermissionToRoleOnDomain)
	r.DELETE("/role/per", handlers.DeletePermissionOfRoleOnDomain)
	r.DELETE("/role", handlers.DeleteRole)

	//get
	r.GET("/domain/role/user", handlers.GetUsersForRoleInDomain)
	r.GET("/domain/user/role", handlers.GetRolesForUserInDomain)
	r.GET("/domain/user/per", handlers.GetPermissionsForUserInDomain)
	r.GET("/domain/user", handlers.GetAllUsersByDomain)

	return r
}
