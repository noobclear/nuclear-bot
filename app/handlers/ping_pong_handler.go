package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
)

func NewPingPongHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *msgs.Context, w msgs.Writer, m msgs.Message) {
			pingMessage, ok := m.(*msgs.PingMessage)
			if ok {
				w.Write(pingMessage.FormatResponse())
			} else {
				h.Handle(ctx, w, m)
			}
		},
	)
}
