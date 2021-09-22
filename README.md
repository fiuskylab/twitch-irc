# twitch-irc
Package to handle Twitch's IRC connection and message

## Example

Import
- > go get github.com/fiuskylab/twitch-irc

```golang
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
			// Can write messages with
			// client.Write("Hello, World!")
		case kill := <-stop:
			break
		}
	}
}
```
