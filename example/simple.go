package main

import (
	"fmt"

	irc "github.com/fiuskylab/twitch-irc"
)

func main() {
	cfg := irc.Config{
		OAuthToken:  "oauth:p9qj03vwu78rscpoasobjmnrt7xj0x",
		BotUsername: "rafiuskybot",
		Channels:    []string{"rafiusky", "EduardoRFS"},
		MaxTries:    5,
	}

	client, err := irc.NewClient(&cfg)

	if err != nil {
		// handler error
		fmt.Println("Error!", err.Error())
	}

	if err := client.JoinChannel("00bex"); err != nil {
		fmt.Println("Error!", err.Error())
	}

	for {
		select {
		case msg := <-client.Messages:
			fmt.Println(msg)
		}
	}
}
