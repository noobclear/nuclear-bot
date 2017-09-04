package commands

import "github.com/sirupsen/logrus"

type CommandFactory interface {
	GetCommand(s string) Command
}

type NukeCommandFactory struct {
	Intents map[string]Command
}

func (f *NukeCommandFactory) GetCommand(s string) Command {
	c, ok := f.Intents[s]
	if ok {
		return c
	}
	logrus.Infof("Could not find command for: [%s]", s)
	return nil
}

func NewCommandFactory() CommandFactory {
	intentMap := make(map[string]Command)
	intentMap["get_pc_specs"] = NewPCSpecsCommand()

	return &NukeCommandFactory{
		Intents: intentMap,
	}
}
