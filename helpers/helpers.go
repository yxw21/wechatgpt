package helpers

import (
	"encoding/xml"
	"github.com/eatmoreapple/openwechat"
	"github.com/yxw21/chatgpt"
	"wechatgpt/config"
)

type Msg struct {
	XMLName xml.Name `xml:"msg"`
	Content string   `xml:"content,attr"`
}

func GetATName(user *openwechat.User) string {
	if user.DisplayName != "" {
		return user.DisplayName
	}
	return user.NickName
}

func GetVerifyMsg(content string) string {
	var msg Msg
	if err := xml.Unmarshal([]byte(content), &msg); err != nil {
		return ""
	}
	return msg.Content
}

func GetNewChat() *chatgpt.Chat {
	if config.Instance.Username != "" && config.Instance.Password != "" && config.Instance.Key != "" {
		return chatgpt.NewChat(config.Instance.Username, config.Instance.Password, config.Instance.Key)
	}
	return chatgpt.NewChatWithSessionToken(config.Instance.SessionToken)
}
