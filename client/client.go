package client

import (
	"github.com/fiuskylab/twitch-irc/auth"
	"github.com/fiuskylab/twitch-irc/internal"
	"github.com/fiuskylab/twitch-irc/irc/listener"
	"github.com/fiuskylab/twitch-irc/irc/message"
)

type Client struct {
	auth        *auth.Auth
	common      *internal.Common
	ircListener *listener.Listener
	Msg         chan message.Message
}

func NewClient(cfg *internal.Config) *Client {
	common := internal.NewCommon(*cfg)

	client := Client{
		common:      common,
		auth:        auth.NewAuth(common),
		ircListener: listener.NewListener(common),
		Msg:         make(chan message.Message, 100),
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
