package config

import (
	"github.com/yxw21/chatgpt"
	"os"
)

type Config struct {
	SessionToken    string
	Username        string
	Password        string
	FriendAddPolicy string
}

var Instance *Config

var Chats = make(map[string]*chatgpt.Chat)

func init() {
	Instance = &Config{
		SessionToken:    os.Getenv("CHAT_GPT_SESSION_TOKEN"),
		Username:        os.Getenv("CHAT_GPT_USERNAME"),
		Password:        os.Getenv("CHAT_GPT_PASSWORD"),
		FriendAddPolicy: os.Getenv("CHAT_GPT_WECHAT_POLICY"),
	}
}
