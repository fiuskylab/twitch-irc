package message

import (
	"strings"
)

// Message refers to each line read
// in IRC connection
type Message struct {
	// Channel is which channel the
	// message was sent
	Channel string

	// Sender the username of who sent
	// the Message
	Sender string

	// Text is the text sent to IRC
	Text string

	// isPing defines if the message was
	// a PING from Twitch
	IsPing bool

	// isNil
	IsNil bool
}

// Types of messages:
// PING :tmi.twitch.tv
// :tmi.twitch.tv 004 rafiuskybot :-
// :ricardinst!ricardinst@ricardinst.tmi.twitch.tv PRIVMSG #rafiusky :Shizukani shite kudasai!

func ParseLine(line string) (msg Message) {
	l := len(line)

	if l < 5 {
		msg.IsNil = true
		return
	}

	if line[:14] == ":tmi.twitch.tv" {
		msg.IsNil = true
		return
	}

	if line[:4] == "PING" {
		msg.IsPing = true
		return
	}

	exclamationPos := strings.IndexByte(line, byte('!'))

	if exclamationPos == -1 {
		msg.IsNil = true
		return
	}

	msg.Sender = line[1:exclamationPos]

	lenUntilChannel := (3 * len(msg.Sender)) + 27

	channelAndMessage := line[lenUntilChannel:]

	colonPos := strings.IndexByte(channelAndMessage, byte(':'))

	if colonPos == -1 {
		msg.IsNil = true
		return
	}

	msg.Channel = channelAndMessage[:colonPos-1]

	msg.Text = channelAndMessage[colonPos+1:]

	return
}
