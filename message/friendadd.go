package message

import (
	"github.com/eatmoreapple/openwechat"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"wechatgpt/config"
	"wechatgpt/helpers"
)

func FriendAdd(ctx *openwechat.MessageContext) {
	if config.Instance.FriendAddPolicy == "ignore" {
		return
	}
	if config.Instance.FriendAddPolicy == "agree" {
		if _, err := ctx.Message.Agree(); err != nil {
			log.Println(err)
			return
		}
	}
	sli := strings.Split(config.Instance.FriendAddPolicy, ",")
	if len(sli) < 2 {
		return
	}
	verify := false
	action := sli[0]
	filter := strings.TrimSpace(sli[1])
	if strings.HasPrefix(filter, "https://") || strings.HasPrefix(filter, "http://") {
		// webhook
		resp, _ := http.Get(filter + "/" + url.PathEscape(helpers.GetVerifyMsg(ctx.Message.Content)))
		if resp != nil {
			verify = resp.StatusCode == 201
			_ = resp.Body.Close()
		}
	} else {
		// regexp
		reg, _ := regexp.Compile(filter)
		if reg != nil {
			verify = reg.Match([]byte(helpers.GetVerifyMsg(ctx.Message.Content)))
		}
	}
	if action == "agree" && verify {
		if _, err := ctx.Message.Agree(); err != nil {
			log.Println(err)
			return
		}
	}
}
