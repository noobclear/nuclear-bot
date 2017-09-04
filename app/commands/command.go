package commands

import "github.com/noobclear/nuclear-bot/app/msgs"

type Command interface {
	Execute(w msgs.Writer, m msgs.PrivMessage) error
}
