package app

import (
	"github.com/noobclear/nuclear-bot/app/handlers"
	"github.com/noobclear/nuclear-bot/app/messages"
)

type Router interface {
	Route(msg string, ctx *messages.Context)
}

type MessageRouter struct {
	Handlers []handlers.Handler
}

func (mr *MessageRouter) Route(msg string, ctx *messages.Context) {
	m := messages.NewMessage(msg)

	for _, h := range mr.Handlers {
		if h.Handle(ctx, m) {
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