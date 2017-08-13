package messages

import (
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/sirupsen/logrus"
)

type PingPongHandler struct {}

func (pph *PingPongHandler) Handle(ctx *Context, m Message) bool {
	isPing := m.MsgParts[0] == "PING"
	if isPing {
		_, err := ctx.Connection.Write([]byte("PONG " + m.MsgParts[1] + util.CRLF))
		if err != nil {
			logrus.WithError(err).Error("Failed to write PONG")
		}
	}
	return isPing
}

func NewPingPongHandler() Handler {
	return &PingPongHandler{}
}
