package handlers

import "github.com/noobclear/nuclear-bot/app/msgs"

type Handler interface {
	Handle(ctx *msgs.Context, w msgs.Writer, m msgs.Message)
}

type HandlerFunc func(ctx *msgs.Context, w msgs.Writer, m msgs.Message)

func (f HandlerFunc) Handle(ctx *msgs.Context, w msgs.Writer, m msgs.Message) {
	f(ctx, w, m)
}
