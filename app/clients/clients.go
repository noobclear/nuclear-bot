package clients

import "github.com/noobclear/nuclear-bot/app/config"

type Clients struct {
	NLP    NLPClient
	Lights LightsClient
	Twitch TwitchClient
}

func NewClients(c config.BotConfig) *Clients {
	return &Clients{
		NLP:    NewWitClient(c.WitAccessToken),
		Lights: NewLIFXClient(c.LIFXAccessToken),
		Twitch: NewTwitchV5Client(c.TwitchClientID),
	}
}
