package app

import (
	"bufio"
	"net"
	"net/textproto"
	"github.com/sirupsen/logrus"
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/noobclear/nuclear-bot/app/messages"
	"github.com/noobclear/nuclear-bot/app/config"
)

type Starter interface {
	Start()
}

type Bot struct {
	config.BotConfig
	Router
}

func (b *Bot) Start() {
	conn := b.getConnection()
	defer conn.Close()

	r := textproto.NewReader(bufio.NewReader(conn))
	ctx := messages.Context{
		Connection: conn,
		BotUsername: b.BotUsername,
		TargetChannel: b.TargetChannel,
	}

	b.listen(&ctx, r)
}

func (b *Bot) getConnection() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")

	if err != nil {
		panic(err)
	}

	// token, username, and channel
	conn.Write([]byte("PASS " + b.BotConfig.TwitchOAuthToken + util.CRLF))
	conn.Write([]byte("NICK " + b.BotConfig.BotUsername + util.CRLF))
	conn.Write([]byte("JOIN " + b.BotConfig.TargetChannel + util.CRLF))

	logrus.Infof("%s joined %s", b.BotUsername, b.TargetChannel)
	return conn
}

func (b *Bot) listen(ctx *messages.Context, r *textproto.Reader) {
	for {
		msg, err := r.ReadLine()
		if err != nil {
			panic(err)
		}

		logrus.Infof("< %s", msg)
		b.Router.Route(msg, ctx)
	}
}
