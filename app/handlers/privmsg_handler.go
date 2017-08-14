package handlers

import (
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/sirupsen/logrus"
	"github.com/noobclear/nuclear-bot/app/messages"
)

type PrivMsgHandler struct {}

func (pmh *PrivMsgHandler) Handle(ctx *messages.Context, m messages.Message) bool {
	privMessage, ok := m.(*messages.PrivMessage)
	if ok {
		resp := "PRIVMSG " + privMessage.ToChannel + " :" + privMessage.Text
		logrus.Infof("Sent message: %s", resp)
		_, err := ctx.Connection.Write([]byte(resp + util.CRLF))
		if err != nil {
			logrus.WithError(err).Error("Failed to write PRIVMSG")
		}
	}
	return ok
}

func NewPrivMsgHandler() Handler {
	return &PrivMsgHandler{}
}
