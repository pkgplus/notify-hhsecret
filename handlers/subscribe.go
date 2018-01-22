package handlers

import (
	"github.com/bingbaba/hhsecret/pkg/client"
	"github.com/kataras/iris/context"
	"github.com/xuebing1110/notify-inspect/pkg/plugin"
	"net/http"
)

func Subscribe(ctx context.Context) {
	s := new(plugin.Subscribe)
	err := ctx.ReadJSON(s)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "ParseJsonFailed", err.Error())
		return
	}

	uid := s.GetParamValue("uid")
	pwd := s.GetParamValue("pwd")
	if uid == "" || pwd == "" {
		SendResponse(ctx, http.StatusBadRequest, "ParamMiss", "uid and pwd was required")
		return
	}

	err = client.Login(uid, pwd)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "LoginFailed", err.Error())
	} else {
		SendNormalResponse(ctx, nil)
	}
}

func GetSubscribe(ctx context.Context) {
	uid := ctx.Params().Get("uid")
	if uid == "" {
		SendResponse(ctx, http.StatusBadRequest, "ParamMissing", "uid param was required")
		return
	}

	login_data, err := client.GetLoginInfo(uid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "LoginFailed", err.Error())
	} else {
		SendNormalResponse(ctx, login_data)
	}
}
