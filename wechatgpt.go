package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"log"
	"wechatgpt/message"
)

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.OnFriendAdd(message.FriendAdd)
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
		if ctx.Message.IsSendByGroup() {
			message.SendByGroup(ctx.Message)
			return
		}
		message.Other(ctx.Message)
	})
	bot.MessageHandler = openwechat.DispatchMessage(dispatcher)
	bot.UUIDCallback = func(uuid string) {
		url := fmt.Sprintf("https://login.weixin.qq.com/qrcode/%s", uuid)
		fmt.Println(url)
	}
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("wechat.json")
	if err := bot.HotLogin(reloadStorage); err != nil {
		if err = bot.Login(); err != nil {
			log.Fatal(err)
		}
	}
	bot.Block()
}
