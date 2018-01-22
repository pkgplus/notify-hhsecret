package v2

import (
	"github.com/xuebing1110/notify-hhsecret/handlers"
)

func plugin() {
	api.Post("/sub/users", handlers.Subscribe)
	api.Get("/sub/users/:userid", handlers.GetSubscribe)
	api.Post("/sub/records", handlers.RecordNotice)
}
