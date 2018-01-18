package v1

import (
	"github.com/xuebing1110/notify-hhsecret/handlers"
)

func plugin() {
	api.Post("/sub/users", handlers.Subscribe)
	api.Post("/sub/records", handlers.RecordNotice)
}
