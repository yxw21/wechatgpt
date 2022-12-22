package config

import (
	"github.com/yxw21/chatgpt"
	session "github.com/yxw21/chatgpt/session"
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
	Session  *session.Session
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
	if Instance.Username != "" && Instance.Password != "" && Instance.Key != "" {
		Session = session.NewSessionWithCredential(Instance.Username, Instance.Password, Instance.Key).AutoRefresh()
	} else {
		Session = session.NewSessionWithAccessToken(Instance.AccessToken).AutoRefresh()
	}
}
