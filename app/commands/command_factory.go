package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/noobclear/nuclear-bot/app/clients"
)

type CommandFactory interface {
	GetCommand(s string) Command
}

type NukeCommandFactory struct {
	Commands map[string]Command
}

func (f *NukeCommandFactory) GetCommand(s string) Command {
	c, ok := f.Commands[s]
	if ok {
		return c
	}
	logrus.Infof("Could not find command for: [%s]", s)
	return nil
}

func NewCommandFactory(clients *clients.Clients) CommandFactory {
	m := make(map[string]Command)
	m["get_pc_specs"] = NewPCSpecsCommand()

	return &NukeCommandFactory{
		Commands: m,
	}
}
