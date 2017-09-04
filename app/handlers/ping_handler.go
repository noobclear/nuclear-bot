package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
	"github.com/noobclear/nuclear-bot/app/request"
)

func NewPingHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *request.Context, w msgs.Writer, m msgs.Message) {
			pingMessage, ok := m.(*msgs.PingMessage)
			if ok {
				w.Write(pingMessage.FormatResponse())
			} else {
				h.Handle(ctx, w, m)
			}
		},
	)
}
