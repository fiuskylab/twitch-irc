package client

import (
	"github.com/fiuskylab/twitch-irc/auth"
	"github.com/fiuskylab/twitch-irc/internal"
	"github.com/fiuskylab/twitch-irc/irc/listener"
	"github.com/fiuskylab/twitch-irc/irc/message"
	twitchapi "github.com/fiuskylab/twitch-irc/twitchapi/client"
)

type Client struct {
	auth        *auth.Auth
	common      *internal.Common
	ircListener *listener.Listener
	Msg         chan message.Message
	TwitchAPI   *twitchapi.Client
}

func NewClient(cfg *internal.Config) *Client {
	common := internal.NewCommon(*cfg)

	client := Client{
		common:      common,
		auth:        auth.NewAuth(common),
		ircListener: listener.NewListener(common),
		Msg:         make(chan message.Message, 100),
	}

	client.TwitchAPI = &twitchapi.Client{
		Auth:   client.auth,
		Common: common,
	}

	// Don't do this at home
	go func() {
		for {
			msg := <-client.ircListener.Messages
			client.Msg <- msg
		}
	}()

	return &client
}
