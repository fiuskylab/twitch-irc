package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fiuskylab/twitch-irc/auth"
	"github.com/fiuskylab/twitch-irc/internal"
	"github.com/fiuskylab/twitch-irc/twitchapi/entity"
	axios "github.com/vicanso/go-axios"
)

const (
	baseURL = `https://api.twitch.tv/helix`
)

var (
	channelURL       = `/channels`
	searchChannelURL = `/search/channels`
)

type Client struct {
	Auth       *auth.Auth
	Common     *internal.Common
	httpClient *axios.Instance
}

func (c *Client) mountHTTPClient() {
	c.httpClient = axios.NewInstance(&axios.InstanceConfig{
		BaseURL: baseURL,
		Headers: http.Header{
			"Authorization": {"Bearer " + c.Auth.BearerToken},
			"Client-Id":     {c.Common.ClientID},
		},
	})
}

func (c *Client) GetChannelInfo(username string) entity.ChannelInformation {
	ent := entity.ChannelInformation{}
	resp, err := c.httpClient.Get(searchChannelURL, url.Values{
		"query": {username},
	})

	if err != nil {
		c.Common.L.Error(err.Error())
		return ent
	}

	respRaw := map[string][]interface{}{}

	if err = resp.JSON(&respRaw); err != nil {
		c.Common.L.Error(err.Error())
		return ent
	}

	if len(respRaw["data"]) == 0 {
		c.Common.L.Error(fmt.Sprintf(`user "%s" not found`, username))
		return ent
	}

	b, err := json.Marshal(respRaw["data"][0])
	if err != nil {
		c.Common.L.Error(err.Error())
		return ent
	}

	if err := json.Unmarshal(b, &ent); err != nil {
		c.Common.L.Error(err.Error())
		return ent
	}

	return ent
}
