package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/bingbaba/hhsecret"
	"github.com/bingbaba/hhsecret/pkg/client"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/xuebing1110/notify-inspect/pkg/plugin"
)

var (
	memCache *cache.Cache
)

func init() {
	memCache = cache.New(6*time.Hour, 10*time.Minute)
}

func RecordNotice(ctx *gin.Context) {
	p := new(plugin.PluginRecord)
	err := ctx.BindJSON(p)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "ParseJsonFailed", err.Error())
		return
	}

	uid := p.GetParamValue("uid")
	if uid == "" {
		SendResponse(ctx, http.StatusBadRequest, "ParamMiss", "can't found \"uid\" param")
		return
	}

	var values []string

	// signed
	key := uid + "," + p.Id
	signed, found := memCache.Get(key)
	if found && signed.(bool) {
		SendNormalResponse(ctx, values)
		return
	}

	// check
	log.Printf("get %s notice...", uid)
	notice, err := client.IfNotice(uid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "NoticeCheckFailed", err.Error())
		return
	}
	log.Printf("get %s notice %v", uid, notice)
	if notice {
		sign_list, err := client.GetListSign(uid)
		if err != nil {
			SendResponse(ctx, http.StatusInternalServerError, "ListSignFailed", err.Error())
			return
		}

		values = noticeInfo(p, sign_list.Signs)
	} else {
		memCache.Set(key, true, 6*time.Hour)
	}

	SendNormalResponse(ctx, values)
}

func noticeInfo(p *plugin.PluginRecord, signs []*hhsecret.Sign) []string {

	uid := p.GetParamValue("uid")
	tip := p.GetParamValue("tip")
	if tip == "" {
		tip = "打卡提醒"
	}

	if len(signs) > 1 {
		return []string{
			uid,
			"懒人打卡",
			signs[0].Location,
			signs[0].GetMinuteSecode(),
			tip,
		}
	} else {
		return []string{
			uid,
			"懒人打卡",
			"--",
			"--",
			tip,
		}
	}
}
