package config

type BotConfig struct {
	TwitchOAuthToken string `json:"twitch_oauth_token"`
	BotUsername      string `json:"bot_username"`
	TargetChannel    string `json:"target_channel"`
	RateLimit        int    `json:"rate_limit"`
	WitAccessToken   string `json:"wit_access_token"`
	LIFXAccessToken  string `json:"lifx_access_token"`
}
