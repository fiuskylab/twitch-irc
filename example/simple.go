package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	irc "github.com/fiuskylab/twitch-irc"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	cfg := irc.Config{
		OAuthToken:  "oauth:your_token",
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
		case kill := <-stop:
			break
		}
	}
}
