package handlers

import (
	"fmt"
	"github.com/noobclear/nuclear-bot/app/msgs"
)

func NewNLPHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *msgs.Context, w msgs.Writer, m msgs.Message) {
			privMessage, ok := m.(*msgs.PrivMessage)
			if ok {
				if len(privMessage.Text) < 2 {
					privMessage.Text = fmt.Sprintf("skipping short message = %s", privMessage.Text)
					w.Write(privMessage.FormatResponse())
				} else {
					// Try to extract intent from text
					i, err := ctx.Clients.NLP.GetIntent(privMessage.Text)
					if err != nil {
						h.Handle(ctx, w, m)
						return
					}
					privMessage.Text = fmt.Sprintf("%+v", i)
					w.Write(privMessage.FormatResponse())
				}
			} else {
				h.Handle(ctx, w, m)
			}
		},
	)
}
