package handlers

import (
	"github.com/noobclear/nuclear-bot/app/msgs"
	"strings"
	"github.com/sirupsen/logrus"
	"github.com/noobclear/nuclear-bot/app/request"
)

var KnownBots = map[string]struct{} {
	"nightbot":  {},
	"newbclear": {},
}

// Filters out msgs from other bots to avoid recursive messaging behavior
func NewIgnoreBotsHandler(h Handler) Handler {
	return HandlerFunc(
		func(ctx *request.Context, w msgs.Writer, m msgs.Message) {
			privMessage, okCast := m.(*msgs.PrivMessage)
			if okCast {
				fromSelf := privMessage.FromUser == ctx.BotUsername
				_, fromBot := KnownBots[strings.ToLower(privMessage.FromUser)]

				// Ignore self and bots
				if !(fromSelf || fromBot) {
					h.Handle(ctx, w, m)
				} else {
					logrus.Infof("Ignoring message from bot: %s", privMessage.FromUser)
				}
			}
		},
	)
}
