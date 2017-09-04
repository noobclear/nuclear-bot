package handlers

import (
	"fmt"
	"github.com/noobclear/nuclear-bot/app/clients"
	"github.com/noobclear/nuclear-bot/app/msgs"
	"github.com/noobclear/nuclear-bot/app/request"
)

func NewNLPHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *request.Context, w msgs.Writer, m msgs.Message) {
			privMessage, ok := m.(*msgs.PrivMessage)
			if ok {
				if len(privMessage.Text) < 2 {
					privMessage.Text = fmt.Sprintf("skipping short message = %s", privMessage.Text)
					w.Write(privMessage.FormatResponse())
				} else {
					// Try to extract intent from text
					resp, err := ctx.Clients.NLP.GetIntent(privMessage.Text)
					if err != nil {
						h.Handle(ctx, w, m)
						return
					}

					i := extractIntent(resp)
					c := ctx.CommandFactory.GetCommand(i)
					if c != nil {
						c.Execute(w, *privMessage)
					}
				}
			} else {
				h.Handle(ctx, w, m)
			}
		},
	)
}

func extractIntent(resp *clients.GetIntentResponse) string {
	if resp == nil {
		return ""
	}
	intents := resp.Entities.Intents
	intentVal := ""
	if len(intents) > 0 {
		for _, i := range intents {
			if i.Confidence > 0.9 {
				intentVal = i.Value
			}
		}
	}
	return intentVal
}
