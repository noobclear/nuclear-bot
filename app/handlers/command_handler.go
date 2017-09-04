package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
	"github.com/noobclear/nuclear-bot/app/request"
	"github.com/sirupsen/logrus"
	"strings"
)

func NewCommandHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *request.Context, w msgs.Writer, m msgs.Message) {
			privMessage, ok := m.(*msgs.PrivMessage)
			if ok {
				// Currently we ignore commands
				msgParts := strings.Split(privMessage.Text, " ")
				if len(msgParts) > 0 {
					if strings.HasPrefix(msgParts[0], "!") {
						logrus.Infof("Ignoring command %s", msgParts[0])
					} else {
						h.Handle(ctx, w, m)
					}
				}
			}
		},
	)
}
