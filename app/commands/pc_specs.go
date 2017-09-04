package commands

import "github.com/noobclear/nuclear-bot/app/msgs"

type PCSpecsCommand struct{}

func (c *PCSpecsCommand) Execute(w msgs.Writer, m msgs.PrivMessage) error {
	// TODO: set depending on target channel
	m.Text = "https://pcpartpicker.com/user/noobclear/saved/zT99WZ"
	w.Write(m.FormatResponse())
	return nil
}

func NewPCSpecsCommand() Command {
	return &PCSpecsCommand{}
}
