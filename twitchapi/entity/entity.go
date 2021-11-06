package entity

/*
"broadcaster_login": "loserfruit",
"display_name": "Loserfruit",
"game_id": "498000",
"game_name": "House Flipper",
"id": "41245072",
"is_live": false,
"tags_ids": [],
"thumbnail_url": "https://static-cdn.jtvnw.net/jtv_user_pictures/fd17325a-7dc2-46c6-8617-e90ec259501c-profile_image-300x300.png",
"title": "loserfruit",
"started_at": ""
* */

type ChannelInformation struct {
	ID               string `json:"id"`
	BroadcasterLogin string `json:"broadcaster_login"`
	DisplayName      string `json:"display_name"`
	GameName         string `json:"game_name"`
	Title            string `json:"title"`
}
