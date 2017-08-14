package handlers

import "github.com/noobclear/nuclear-bot/app/messages"

type IgnoreSelfHandler struct {}

// Filters out messages from other instances of the same bot to avoid recursive messaging behavior
func (ish *IgnoreSelfHandler) Handle(ctx *messages.Context, m messages.Message) bool {
	privMessage, ok := m.(*messages.PrivMessage)
	return ok && privMessage.FromUser == ctx.BotUsername
}

func NewIgnoreSelfHandler() Handler {
	return &IgnoreSelfHandler{}
}