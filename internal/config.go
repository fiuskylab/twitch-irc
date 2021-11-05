package internal

// Config store the base information
// for a connection with Twitch's IRC.
type Config struct {
	// OAuthToken is the token to allow you
	// to connect to the IRC.
	OAuthToken string

	// BotUsername must be the same as the account
	// that was used to generate the OAuthToken.
	BotUsername string

	// Channels are the names of a each Twitch's
	// channel that you want to your bot to connect to.
	Channels []string

	// MaxTries set de amount of times between
	// each connection retry.
	MaxTries uint

	// ClientID is the APP id from
	// https://dev.twitch.tv/console/apps
	ClientID string

	// ClientSecret is the APP secret generated
	// on https://dev.twitch.tv/console/apps
	ClientSecret string
}
