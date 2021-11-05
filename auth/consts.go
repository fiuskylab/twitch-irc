package auth

const (
	daySeconds = 60 * 60 * 24
)

var (
	scope = []string{
		"analytics:read:extensions",
		"analytics:read:games",
		"bits:read",
		"channel:edit:commercial",
		"channel:manage:broadcast",
		"channel:manage:extensions",
		"channel:manage:polls",
		"channel:manage:predictions",
		"channel:manage:redemptions",
		"channel:manage:schedule",
		"channel:manage:videos",
		"channel:moderate",
		"channel:read:editors",
		"channel:read:goals",
		"channel:read:hype_train",
		"channel:read:polls",
		"channel:read:predictions",
		"channel:read:redemptions",
		"channel:read:stream_key",
		"channel:read:subscriptions",
		"chat:edit",
		"chat:read",
		"clips:edit",
		"moderation:read",
		"moderator:manage:automod",
		"user:edit",
		"user:edit:follows",
		"user:manage:blocked_users",
		"user:read:blocked_users",
		"user:read:broadcast",
		"user:read:email",
		"user:read:follows",
		"user:read:subscriptions",
		"whispers:read",
		"whispers:edit",
	}
)
