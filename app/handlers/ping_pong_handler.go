package handlers

import (
	"github.com/noobclear/nuclear-bot/app/messages"
)

type PingPongHandler struct{}

func (pph *PingPongHandler) Handle(ctx *messages.Context, m messages.Message) (string, bool, error) {
	pingMessage, ok := m.(*messages.PingMessage)
	if ok {
		resp := pingMessage.FormatResponse()
		return resp, true, nil
	}
	return "", false, nil
}

func NewPingPongHandler() Handler {
	return &PingPongHandler{}
}
