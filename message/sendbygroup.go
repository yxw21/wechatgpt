package message

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/yxw21/chatgpt"
	"strings"
	"wechatgpt/config"
	"wechatgpt/helpers"
)

func SendByGroup(msg *openwechat.Message) {
	var word string
	if !msg.IsAt() {
		return
	}
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		_, _ = msg.ReplyText(err.Error())
		return
	}
	sender, err := msg.Sender()
	if err != nil {
		_, _ = msg.ReplyText("@" + helpers.GetATName(groupSender) + "\n" + err.Error())
		return
	}
	senderID := sender.ID()
	if _, ok := config.Chats[senderID]; !ok {
		config.Chats[senderID] = chatgpt.NewChat(config.Session)
	}
	word = strings.Replace(msg.Content, "@"+sender.Self.NickName, "", -1)
	for i := 0; i < config.Instance.MsgRetry; i++ {
		res, err := config.Chats[senderID].Send(word)
		if err == nil {
			_, _ = msg.ReplyText("@" + helpers.GetATName(groupSender) + "\n" + res.Message.Content.Parts[0])
			break
		}
		if i == config.Instance.MsgRetry-1 {
			_, _ = msg.ReplyText("@" + helpers.GetATName(groupSender) + "\n" + err.Error())
		}
	}
}
