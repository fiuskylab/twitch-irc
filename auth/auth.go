package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/fiuskylab/twitch-irc/internal"
)

type TwitchResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_id"`
	TokenType   string `json:"token_type"`
}

type Auth struct {
	common      *internal.Common
	BearerToken string
}

func NewAuth(c *internal.Common) *Auth {

	a := Auth{
		common: c,
	}

	if err := a.setOAuth(); err != nil {
		a.common.L.Error(err.Error())
	}

	go func() {
		time.Sleep(time.Second * daySeconds)
		c.setOAuth()
	}()

	return c
}

const (
	twitch_base_url = `https://id.twitch.tv/oauth2/token`
)

func (a *Auth) setOAuth() error {
	formVals := url.Values{
		"client_id":     {a.common.CLIENT_ID},
		"client_secret": {a.common.CLIENT_SECRET},
		"grant_type":    {"client_credentials"},
		"scope":         {strings.Join(scope, " ")},
	}

	resp, err := http.PostForm(twitch_base_url, formVals)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	twitchResp := TwitchResp{}

	if err := json.Unmarshal(body, &twitchResp); err != nil {
		return err
	}

	a.BearerToken = twitchResp.AccessToken

	return nil
}
