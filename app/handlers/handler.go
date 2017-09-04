package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
	"github.com/noobclear/nuclear-bot/app/request"
)

type Handler interface {
	Handle(ctx *request.Context, w msgs.Writer, m msgs.Message)
}

type HandlerFunc func(ctx *request.Context, w msgs.Writer, m msgs.Message)

func (f HandlerFunc) Handle(ctx *request.Context, w msgs.Writer, m msgs.Message) {
	f(ctx, w, m)
}
