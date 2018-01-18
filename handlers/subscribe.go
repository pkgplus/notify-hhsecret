package handlers

import (
	"github.com/bingbaba/hhsecret/pkg/client"
	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify-inspect/pkg/plugin"
	"net/http"
)

func Subscribe(ctx context.Context) {
	p := new(plugin.Plugin)
	err := ctx.ReadJSON(p)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "ParseJsonFailed", err.Error())
		return
	}

	uid := p.GetParamValue("uid")
	pwd := p.GetParamValue("pwd")
	if uid == "" || pwd == "" {
		SendResponse(ctx, http.StatusBadRequest, "ParamMiss", "uid and pwd was required")
		return
	}

	login_data, err := client.Login(uid, pwd)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "LoginFailed", err.Error())
	} else {
		SendNormalResponse(ctx, login_data)
	}
}
