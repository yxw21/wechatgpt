package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
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
		qr, err := qrcode.New(fmt.Sprintf("https://login.weixin.qq.com/l/%s", uuid), qrcode.Medium)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(qr.ToSmallString(true))
		}
	}
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("wechat.json")
	if err := bot.HotLogin(reloadStorage); err != nil {
		if err = bot.Login(); err != nil {
			log.Fatal(err)
		}
	}
	bot.Block()
}
