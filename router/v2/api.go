package v2

import (
	"github.com/kataras/iris/core/router"
	"github.com/xuebing1110/notify-hhsecret/app"
)

var api router.Party

func init() {
	irisApp := app.GetIrisApp()
	api = irisApp.Party("/api/v2/plugins/" + app.APP_NAME)

	plugin()
}
