package msgs

import "github.com/noobclear/nuclear-bot/app/clients"

type Context struct {
	BotUsername   string
	TargetChannel string
	Clients       *clients.Clients
}
