package v2

import (
	"github.com/xuebing1110/notify-hhsecret/handlers"
)

func init() {
	api.POST("/sub/users", handlers.Subscribe)
	api.GET("/sub/users/:uid", handlers.GetSubscribe)
	api.POST("/sub/records", handlers.RecordNotice)
}
