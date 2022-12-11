package config

import (
	"github.com/yxw21/chatgpt"
	"os"
)

type Config struct {
	Username        string
	Password        string
	SessionToken    string
	FriendAddPolicy string
}

var Instance *Config

var Chats = make(map[string]*chatgpt.Chat)

func init() {
	Instance = &Config{
		Username:        os.Getenv("WECHAT_CHAT_GPT_USERNAME"),
		Password:        os.Getenv("WECHAT_CHAT_GPT_PASSWORD"),
		SessionToken:    os.Getenv("WECHAT_CHAT_GPT_TOKEN"),
		FriendAddPolicy: os.Getenv("WECHAT_CHAT_GPT_POLICY"),
	}
}
