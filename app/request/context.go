package request

import (
	"github.com/noobclear/nuclear-bot/app/clients"
	"github.com/noobclear/nuclear-bot/app/commands"
)

type Context struct {
	BotUsername    string
	TargetChannel  *clients.User
	Clients        *clients.Clients
	CommandFactory commands.CommandFactory
}
