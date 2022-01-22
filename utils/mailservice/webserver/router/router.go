package router

import (
	"fmt"
	"time"

	"github.com/vfluxus/mailservice/webserver/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vfluxus/mailservice/core"
	_ "github.com/vfluxus/mailservice/docs" // import for swagger generated docs
	"github.com/vfluxus/mailservice/webserver/controller"
)

// ------------------------------
// NewRouter ...
// @title Mail service API
// @version 1.0
// @description Mail service api.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:10000
// @BasePath
func NewRouter() (routers *gin.Engine) {
	routers = gin.New()
	routers.Use(gin.Recovery())
	routers.Use(gin.Logger())

	//CORS
	routers.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// swager
	// The url pointing to API definition
	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/docs/doc.json",
		core.GetConfig().Web.Host, core.GetConfig().Web.Port))
	routers.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// authentication
	routers.Use(middlewares.DecodeToken())

	accGr := routers.Group("/account")
	{
		accCtrl := new(controller.AccountController)
		accGr.GET("", accCtrl.GetAccount)
		accGr.POST("", accCtrl.CreateNewAccount)
		accGr.PUT("", accCtrl.UpdateAccountInfo)
		accGr.DELETE("", accCtrl.DeleteAccount)
	}

	templateGr := routers.Group("/template")
	{
		templateCtrl := new(controller.TemplateController)
		templateGr.GET("", templateCtrl.GetTemplate)
		templateGr.POST("", templateCtrl.CreateNewTemplate)
		templateGr.PUT("", templateCtrl.UpdateTemplate)
		templateGr.DELETE("", templateCtrl.DeleteTemplate)
	}

	mailGr := routers.Group("/mail")
	{
		mailCtrl := new(controller.MailController)
		mailGr.GET("", mailCtrl.GetMail)
		mailGr.POST("", mailCtrl.CreateNewMail)
		mailGr.POST("/send", mailCtrl.SendMail)
		mailGr.PUT("", mailCtrl.UpdateMail)
		mailGr.DELETE("", mailCtrl.DeleteMail)
	}

	return routers
}
