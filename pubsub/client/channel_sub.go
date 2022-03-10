package client

func (c *Client) Subscribe(username string) {
	channelID := c.TwitchAPI.GetChannelInfo(username).ID

}
