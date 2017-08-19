package handlers

import (
	"github.com/noobclear/nuclear-bot/app/messages"
)

type PrivMsgHandler struct{}

func (pmh *PrivMsgHandler) Handle(ctx *messages.Context, m messages.Message) (string, bool, error) {
	privMessage, ok := m.(*messages.PrivMessage)
	if ok {
		resp := privMessage.FormatResponse()
		return resp, true, nil
	}
	return "", false, nil
}

func NewPrivMsgHandler() Handler {
	return &PrivMsgHandler{}
}
