package config

import (
	"github.com/yxw21/chatgpt"
	"os"
)

type Config struct {
	Username        string
	Password        string
	Key             string
	AccessToken     string
	FriendAddPolicy string
}

var (
	Instance *Config
	Session  *chatgpt.Session
	Browser  *chatgpt.Browser
	Chats    = make(map[string]*chatgpt.Chat)
)

func init() {
	Instance = &Config{
		Username:        os.Getenv("WECHAT_CHAT_GPT_USERNAME"),
		Password:        os.Getenv("WECHAT_CHAT_GPT_PASSWORD"),
		Key:             os.Getenv("WECHAT_KEY"),
		AccessToken:     os.Getenv("WECHAT_CHAT_GPT_ACCESS_TOKEN"),
		FriendAddPolicy: os.Getenv("WECHAT_CHAT_GPT_POLICY"),
	}
}
