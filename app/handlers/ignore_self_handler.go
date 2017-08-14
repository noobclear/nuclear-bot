package handlers

import (
	"github.com/noobclear/nuclear-bot/app/messages"
)

type IgnoreSelfHandler struct{}

// Filters out messages from other instances of the same bot to avoid recursive messaging behavior
func (ish *IgnoreSelfHandler) Handle(ctx *messages.Context, m messages.Message) (string, bool, error) {
	privMessage, ok := m.(*messages.PrivMessage)
	if ok && privMessage.FromUser == ctx.BotUsername {
		return "", true, nil
	} else {
		return "", false, nil
	}
}

func NewIgnoreSelfHandler() Handler {
	return &IgnoreSelfHandler{}
}
