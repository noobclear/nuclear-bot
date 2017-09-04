package commands

import "github.com/sirupsen/logrus"

type Commander interface {
	Command(s string) Command
}

type CommandFactory struct {
	Intents map[string]Command
}

func (f *CommandFactory) Command(s string) Command {
	c, ok := f.Intents[s]
	if ok {
		return c
	}
	logrus.Infof("Could not find command for: [%s]", s)
	return nil
}

func NewCommandFactory() Commander {
	intentMap := make(map[string]Command)
	intentMap["get_pc_specs"] = NewPCSpecsCommand()

	return &CommandFactory{
		Intents: intentMap,
	}
}

