package app

import (
	"github.com/noobclear/nuclear-bot/app/handlers"
	"github.com/noobclear/nuclear-bot/app/messages"
	"github.com/sirupsen/logrus"
)

type Router interface {
	Route(ctx *messages.Context, msg string, q chan<- string)
}

type MessageRouter struct {
	Handlers []handlers.Handler
}

// Route a message to the handlers and the bot response to the channel
func (mr *MessageRouter) Route(ctx *messages.Context, msg string, q chan<- string) {
	m := messages.NewMessage(msg)

	for _, h := range mr.Handlers {
		resp, f, err := h.Handle(ctx, m)
		if err != nil {
			logrus.WithError(err).Errorf("Encountered error in handler: %T", h)
			return
		}

		if resp != "" {
			q <- resp
		}

		if f {
			return
		}
	}
}

func (mr *MessageRouter) RegisterHandler(h handlers.Handler) {
	mr.Handlers = append(mr.Handlers, h)
}

func NewMessageRouter() Router {
	mr := MessageRouter{}
	mr.RegisterHandler(handlers.NewPingPongHandler())
	mr.RegisterHandler(handlers.NewIgnoreSelfHandler())
	mr.RegisterHandler(handlers.NewPrivMsgHandler())
	return &mr
}