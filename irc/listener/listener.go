package listener

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"strings"
	"time"

	"github.com/fiuskylab/twitch-irc/helper"
	"github.com/fiuskylab/twitch-irc/internal"
	"github.com/fiuskylab/twitch-irc/irc/message"
)

type Listener struct {
	common   *internal.Common
	tp       *textproto.Reader
	conn     net.Conn
	Messages chan message.Message
}

const (
	IRC_URL = "irc.chat.twitch.tv:6667"
)

func NewListener(c *internal.Common) *Listener {
	l := Listener{
		common:   c,
		Messages: make(chan message.Message, 100),
	}

	return &l
}

// AddChannel connects the bot to a new channel
func (l *Listener) JoinChannel(name string) error {
	name = strings.ToLower(name)
	if _, ok := helper.InArrayStr(l.common.Channels, name); ok {
		return nil
	}
	if err := l.Write(string("JOIN #" + name)); err != nil {
		return err
	}
	l.common.Channels = append(l.common.Channels, name)
	return nil
}

// PartChannel leaves a Twitch's channel
func (l *Listener) PartChannel(name string) error {
	name = strings.ToLower(name)
	pos, ok := helper.InArrayStr(l.common.Channels, name)
	if !ok {
		return nil
	}
	l.common.Channels = append(l.common.Channels[:pos], l.common.Channels[pos+1:]...)
	return l.Write(string("PART #" + name))
}

func (l *Listener) setConnection() error {
	if err := l.setTCPConn(); err != nil {
		return err
	}

	if err := l.connectIRC(); err != nil {
		return err
	}

	if err := l.joinChannels(); err != nil {
		return err
	}
	l.setTPReader()
	return nil
}

func (l *Listener) joinChannels() error {
	for _, ch := range l.common.Channels {
		if err := l.JoinChannel(ch); err != nil {
			return err
		}
	}
	return nil
}

// setTCPConn dials and set TCP connection
func (l *Listener) setTCPConn() error {
	var conn net.Conn
	var err error

	for i := uint(0); i < l.common.MaxTries; i++ {
		conn, err = net.Dial("tcp", IRC_URL)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			l.conn = conn
		}
	}
	return err
}

func (l *Listener) connectIRC() error {
	if err := l.
		Write(string("PASS " + l.common.OAuthToken)); err != nil {
		return err
	}
	if err := l.
		Write(string("NICK " + l.common.BotUsername)); err != nil {
		return err
	}
	return nil
}

func (l *Listener) setTPReader() {
	reader := bufio.NewReader(l.conn)
	l.tp = textproto.NewReader(reader)
}

// Write receives a string and write it
// into IRC TCP connection, don't need
// to add "\r\n" at the end of the string.
func (l *Listener) Write(msg string) error {
	ln := len(msg)

	if ln < 3 {
		return fmt.Errorf("Message with lenght < 3")
	}

	if msg[ln-2:] != "\r\n" {
		msg = string(msg + "\r\n")
	}

	_, err := fmt.Fprint(l.conn, msg)

	return err
}

func (l *Listener) listen() {
	for {
		ircLine, err := l.tp.ReadLine()
		if err != nil {
			if err := l.setConnection(); err != nil {
				break
			}
		} else {
			msg := message.ParseLine(ircLine)
			switch {
			case msg.IsNil:
				continue
			case msg.IsPing:
				l.Write("PONG")
			default:
				l.Messages <- msg
			}
		}
	}
}
