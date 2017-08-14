package handlers

import "github.com/noobclear/nuclear-bot/app/messages"

type Handler interface {
	Handle(ctx *messages.Context, m messages.Message) (success bool)
}