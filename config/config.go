package config

import (
	"github.com/yxw21/chatgpt"
	"os"
)

type Config struct {
	SessionToken    string
	FriendAddPolicy string
}

var Instance *Config

var Chats = make(map[string]*chatgpt.Chat)

func init() {
	Instance = &Config{
		SessionToken:    os.Getenv("CHAT_GPT_SESSION_TOKEN"),
		FriendAddPolicy: os.Getenv("CHAT_GPT_WECHAT_POLICY"),
	}
}
