package app

import (
	// "os"

	"github.com/kataras/iris"
	"github.com/xuebing1110/notify-inspect/pkg/plugin"
	"github.com/xuebing1110/notify-inspect/pkg/plugin/client"
)

var (
	APP_NAME string
	IrisApp  *iris.Application
)

func init() {
	// // app name
	// APP_NAME = os.Getenv("APP_NAME")
	// if APP_NAME == "" {
	// 	panic("the env \"APP_NAME\" is required!")
	// }
	APP_NAME = "hhsecret"

	// register plugin
	err := register()
	if err != nil {
		panic("register plugin failed: " + err.Error())
	}

	IrisApp = iris.New()
}

func register() error {
	c := client.DefaultRegisterClient
	p := &plugin.Plugin{
		Id:            APP_NAME,
		Description:   "懒人打卡",
		ServeAddr:     "http://127.0.0.1:8081/api/v1/plugins",
		TemplateMsgId: "8U98v1g7PWLZ5p4jbWNSpY5dr-hhG5kVuMAUew4PHnY",
		Params: []plugin.PluginParam{
			plugin.PluginParam{
				Id:    "uid",
				Name:  "工号",
				Value: "",
			},
			plugin.PluginParam{
				Id:    "pwd",
				Name:  "密码",
				Value: "",
			},
		},
		RecordParams: []plugin.PluginParam{
			plugin.PluginParam{
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

func GetIrisApp() *iris.Application {
	return IrisApp
}
