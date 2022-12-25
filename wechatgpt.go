package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
	"github.com/yxw21/chatgpt"
	"log"
	"wechatgpt/config"
	"wechatgpt/message"
)

func main() {
	browser, closeBrowser, err := chatgpt.NewBrowser(config.Instance.Key)
	if err != nil {
		log.Fatal(err)
	}
	defer closeBrowser()
	config.Browser = browser
	if config.Instance.Username != "" && config.Instance.Password != "" {
		config.Session = chatgpt.NewSessionWithCredential(browser, config.Instance.Username, config.Instance.Password).AutoRefresh()
	} else {
		config.Session = chatgpt.NewSessionWithAccessToken(browser, config.Instance.AccessToken).AutoRefresh()
	}
	bot := openwechat.DefaultBot(openwechat.Desktop)
	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.OnFriendAdd(message.FriendAdd)
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
		if ctx.Message.IsSendByGroup() {
			go message.SendByGroup(ctx.Message)
			return
		}
		go message.Other(ctx.Message)
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
	_ = bot.Block()
}
