package main

import (
	"github.com/xuebing1110/notify-hhsecret/app"
	"github.com/xuebing1110/notify-inspect/pkg/plugin"
	"github.com/xuebing1110/notify-inspect/pkg/plugin/client"

	_ "github.com/xuebing1110/notify-hhsecret/router/v2"
)

func main() {
	app.GetApp().Run()
}

func init() {
	// register plugin
	err := register()
	if err != nil {
		panic("register plugin failed: " + err.Error())
	}
}

func register() error {
	c := client.DefaultRegisterClient
	p := &plugin.Plugin{
		Id:            app.APP_NAME,
		Description:   "懒人打卡",
		ServerAddr:    "https://m.bingbaba.com/api/v2/plugins/" + app.APP_NAME,
		TemplateMsgId: "8U98v1g7PWLZ5p4jbWNSpY5dr-hhG5kVuMAUew4PHnY",
		Emphasis:      "5",
		Params: []plugin.PluginParam{
			{
				Id:    "uid",
				Name:  "工号",
				Value: "",
			},
			{
				Id:    "pwd",
				Name:  "密码",
				Value: "",
			},
		},
		RecordParams: []plugin.PluginParam{
			{
				Id:         "tip",
				Name:       "提示语",
				Value:      "",
				Candidates: []plugin.PluginData{},
			},
		},
		Author: "bingbaba.com",
	}

	return c.Register(p)
}
