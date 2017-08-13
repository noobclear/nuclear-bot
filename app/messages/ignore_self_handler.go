package messages

import (
	"strings"
	"fmt"
)

type IgnoreSelfHandler struct {}

// Attempts to avoid recursive bot behavior in case there is more than 1 bot instance running
func (ish *IgnoreSelfHandler) Handle(ctx *Context, m Message) bool {
	return len(m.MsgParts) > 1 && m.MsgParts[1] == "PRIVMSG" &&
		strings.Split(m.MsgParts[0], "!")[0] == fmt.Sprintf(":%s", ctx.BotUsername)
}

func NewIgnoreSelfHandler() Handler {
	return &IgnoreSelfHandler{}
}