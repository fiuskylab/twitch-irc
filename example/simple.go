package main

import (
	"fmt"

	irc "github.com/fiuskylab/twitch-irc"
)

func main() {
	cfg := irc.Config{
		OAuthToken:  "oauth:dskasdjkasd983819213nsdhj",
		BotUsername: "bot_name",
		Channel:     "channel_name",
		MaxTries:    5,
	}

	client, err := irc.NewClient(&cfg)

	if err != nil {
		// handler error
		fmt.Println("Error!", err.Error())
	}

	for {
		select {
		case msg := <-client.Messages:
			fmt.Println(msg)
		}
	}
}
