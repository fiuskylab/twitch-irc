package twitchirc

import "net/textproto"

// Client is the IRC connector.
type Client struct {
	cfg *Config
	tp  *textproto.Reader
}

// NewClient returns a Client instance
func NewClient(cfg *Config) (*Client, error) {
	c := Client{cfg: cfg}

	return &c, nil
}
