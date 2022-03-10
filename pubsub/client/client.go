package client

import (
	"github.com/fiuskylab/twitch-irc/auth"
	"github.com/fiuskylab/twitch-irc/eventsub/entity"
	"github.com/fiuskylab/twitch-irc/internal"
	"github.com/fiuskylab/twitch-irc/twitchapi/client"
)

type Client struct {
	Auth      *auth.Auth
	Common    *internal.Common
	TwitchAPI *client.Client
	Events    chan entity.Event
}
