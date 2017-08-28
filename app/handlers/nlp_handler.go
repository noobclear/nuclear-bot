package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
	"fmt"
)

func NewNLPHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *msgs.Context, w msgs.Writer, m msgs.Message) {
			privMessage, ok := m.(*msgs.PrivMessage)
			if ok {
				privMessage.Text = fmt.Sprintf("bot received msg = %s", privMessage.Text)
				w.Write(privMessage.FormatResponse())
			} else {
				h.Handle(ctx, w, m)
			}
		},
	)
}
