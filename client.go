package twitchirc

import (
	"fmt"
	"net"
	"net/textproto"
	"time"
)

// Client is the IRC connector.
type Client struct {
	cfg  *Config
	tp   *textproto.Reader
	conn *net.Conn
}

const (
	IRC_URL = "irc://irc.chat.twitch.tv:6667"
)

// NewClient returns a Client instance
func NewClient(cfg *Config) (*Client, error) {
	c := Client{cfg: cfg}

	if err := c.setTCPConn(); err != nil {
		return c, err
	}

	return &c, nil
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
			c.conn = &conn
		}
	}
	return err
}
func (c *Client) Write(msg string) error {
	msg = string(msg + "\r\n")
	_, err := fmt.Fprint(c.conn, msg)

	return err
}
