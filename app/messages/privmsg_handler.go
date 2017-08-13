package messages

import (
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/sirupsen/logrus"
)

type PrivMsgHandler struct {}

func (pmh *PrivMsgHandler) Handle(ctx *Context, m Message) bool {
	isPrivMsg := len(m.MsgParts) > 1 && m.MsgParts[1] == "PRIVMSG"
	if isPrivMsg {
		_, err := ctx.Connection.Write([]byte("PRIVMSG " + m.MsgParts[2] + " " + m.MsgParts[3] + util.CRLF))
		if err != nil {
			logrus.WithError(err).Error("Failed to write PRIVMSG")
		}
	}
	return isPrivMsg
}

func NewPrivMsgHandler() Handler {
	return &PrivMsgHandler{}
}
