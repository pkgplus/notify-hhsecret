package v2

import (
	"github.com/xuebing1110/notify-hhsecret/app"
)

var (
	api = app.GetApp().Group("/api/v2/plugins/" + app.APP_NAME)
)
