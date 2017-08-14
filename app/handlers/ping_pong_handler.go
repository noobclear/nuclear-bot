package handlers

import (
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/sirupsen/logrus"
	"github.com/noobclear/nuclear-bot/app/messages"
)

type PingPongHandler struct {}

func (pph *PingPongHandler) Handle(ctx *messages.Context, m messages.Message) bool {
	pingMessage, ok := m.(*messages.PingMessage)
	if ok {
		_, err := ctx.Connection.Write([]byte("PONG " + pingMessage.Host + util.CRLF))
		if err != nil {
			logrus.WithError(err).Error("Failed to write PONG")
		}
	}
	return ok
}

func NewPingPongHandler() Handler {
	return &PingPongHandler{}
}