package message

import (
	"github.com/eatmoreapple/openwechat"
	"wechatgpt/config"
	"wechatgpt/helpers"
)

func Other(msg *openwechat.Message) {
	sender, err := msg.Sender()
	if err != nil {
		_, _ = msg.ReplyText(err.Error())
		return
	}
	senderID := sender.ID()
	if msg.IsText() {
		if _, ok := config.Chats[senderID]; !ok {
			config.Chats[senderID] = helpers.GetNewChat()
		}
		res, err := config.Chats[senderID].Send(msg.Content)
		if err != nil {
			_, _ = msg.ReplyText(err.Error())
		} else {
			_, _ = msg.ReplyText(res.Message.Content.Parts[0])
		}
	} else {
		_, _ = msg.ReplyText("机器人暂时只支持处理文本消息")
	}
}
