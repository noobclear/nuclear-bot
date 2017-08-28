package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
)

// Filters out msgs from other instances of the same bot to avoid recursive messaging behavior
func NewIgnoreSelfHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *msgs.Context, w msgs.Writer, m msgs.Message) {
			privMessage, ok := m.(*msgs.PrivMessage)
			fromSelf := ok && privMessage.FromUser == ctx.BotUsername

			if !fromSelf {
				h.Handle(ctx, w, m)
			}
		},
	)
}
