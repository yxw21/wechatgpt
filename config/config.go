package config

import (
	"github.com/yxw21/chatgpt"
	"os"
	"strconv"
)

type Config struct {
	Username        string
	Password        string
	Key             string
	AccessToken     string
	FriendAddPolicy string
	MsgRetry        int
}

var (
	Instance *Config
	Session  *chatgpt.Session
	Chats    = make(map[string]*chatgpt.Chat)
)

func init() {
	Instance = &Config{
		Username:        os.Getenv("WECHAT_CHAT_GPT_USERNAME"),
		Password:        os.Getenv("WECHAT_CHAT_GPT_PASSWORD"),
		Key:             os.Getenv("WECHAT_KEY"),
		AccessToken:     os.Getenv("WECHAT_CHAT_GPT_ACCESS_TOKEN"),
		FriendAddPolicy: os.Getenv("WECHAT_CHAT_GPT_POLICY"),
		MsgRetry:        3,
	}
	if Instance.Username != "" && Instance.Password != "" && Instance.Key != "" {
		Session = chatgpt.NewSessionWithCredential(Instance.Username, Instance.Password, Instance.Key).AutoRefresh()
	} else {
		Session = chatgpt.NewSessionWithAccessToken(Instance.AccessToken).AutoRefresh()
	}
	msgRetry, err := strconv.Atoi(os.Getenv("WECHAT_MSG_RETRY"))
	if err == nil && msgRetry > 0 {
		Instance.MsgRetry = msgRetry
	}
}
