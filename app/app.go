package app

import (
	// "os"

	"github.com/gin-gonic/gin"
)

var (
	APP_NAME string
	app      *gin.Engine
)

func init() {
	// // app name
	// APP_NAME = os.Getenv("APP_NAME")
	// if APP_NAME == "" {
	// 	panic("the env \"APP_NAME\" is required!")
	// }
	APP_NAME = "hhsecret"

	app = gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
}

func GetApp() *gin.Engine {
	return app
}
