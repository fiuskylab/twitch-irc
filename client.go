package twitchirc

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"time"
)

// Client is the IRC connector.
type Client struct {
	cfg      *Config
	tp       *textproto.Reader
	conn     net.Conn
	Messages chan Message
}

const (
	IRC_URL = "irc.chat.twitch.tv:6667"
)

// NewClient returns a Client instance
func NewClient(cfg *Config) (*Client, error) {
	c := Client{
		cfg:      cfg,
		Messages: make(chan Message, 100),
	}

	if err := c.setConnection(); err != nil {
		return &c, err
	}

	go c.listen()

	return &c, nil
}

func (c *Client) setConnection() error {
	if err := c.setTCPConn(); err != nil {
		return err
	}

	if err := c.connectIRC(); err != nil {
		return err
	}
	c.setTPReader()
	return nil
}

// setTCPConn dials and set TCP connection
func (c *Client) setTCPConn() error {
	var conn net.Conn
	var err error

	for i := uint(0); i < c.cfg.MaxTries; i++ {
		conn, err = net.Dial("tcp", IRC_URL)
		if err != nil {
			time.Sleep(time.Second)
		} else {
			c.conn = conn
		}
	}
	return err
}

func (c *Client) connectIRC() error {
	if err := c.
		Write(string("PASS " + c.cfg.OAuthToken)); err != nil {
		return err
	}
	if err := c.
		Write(string("NICK " + c.cfg.BotUsername)); err != nil {
		return err
	}
	if err := c.
		Write(string("JOIN #" + c.cfg.Channel)); err != nil {
		return err
	}
	return nil
}

func (c *Client) setTPReader() {
	reader := bufio.NewReader(c.conn)
	c.tp = textproto.NewReader(reader)
}

// Write receives a string and write it
// into IRC TCP connection, don't need
// to add "\r\n" at the end of the string.
func (c *Client) Write(msg string) error {
	l := len(msg)

	if l < 3 {
		return fmt.Errorf("Message with lenght < 3")
	}

	if msg[l-2:] != "\r\n" {
		msg = string(msg + "\r\n")
	}

	_, err := fmt.Fprint(c.conn, msg)

	return err
}

func (c *Client) listen() {
	for {
		ircLine, err := c.tp.ReadLine()
		if err != nil {
			if err := c.setConnection(); err != nil {
				break
			}
		} else {
			msg := parseLine(ircLine)
			switch {
			case msg.isNil:
				continue
			case msg.isPing:
				c.Write("PONG")
			default:
				c.Messages <- msg
			}
		}
	}
}
