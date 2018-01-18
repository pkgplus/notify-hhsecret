package main

import (
	"os"

	"github.com/kataras/iris"
	"github.com/xuebing1110/notify-hhsecret/app"

	_ "github.com/xuebing1110/notify-hhsecret/router/v1"
)

func main() {
	// http server
	irisApp := app.GetIrisApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	irisApp.Run(iris.Addr(":" + port))
}
