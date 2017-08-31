package clients

import "github.com/noobclear/nuclear-bot/app/config"

type Clients struct {
	NLP NLPClient
}

func NewClients(c config.BotConfig) *Clients {
	return &Clients{
		NLP: NewWitClient(c.WitAccessToken),
	}
}
