package twitchirc

// Config store the base information
// for a connection with Twitch's IRC.
type Config struct {
	// OAuthToken is the token to allow you
	// to connect to the IRC.
	OAuthToken string

	// BotUsername must be the same as the account
	// that was used to generate the OAuthToken.
	BotUsername string

	// Channel is the name of a Twitch's channel
	// that you want to your bot to connect to.
	Channel string

	// MaxTries set de amount of times between
	// each connection retry.
	MaxTries uint
}
