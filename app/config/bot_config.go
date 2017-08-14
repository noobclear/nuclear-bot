package config

type BotConfig struct {
	TwitchOAuthToken string `json:"twitch_oauth_token"`
	BotUsername      string `json:"bot_username"`
	TargetChannel    string `json:"target_channel"`
}
