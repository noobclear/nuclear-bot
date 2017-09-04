package app

import (
	"bufio"
	"github.com/noobclear/nuclear-bot/app/clients"
	"github.com/noobclear/nuclear-bot/app/commands"
	"github.com/noobclear/nuclear-bot/app/config"
	"github.com/noobclear/nuclear-bot/app/handlers"
	"github.com/noobclear/nuclear-bot/app/msgs"
	"github.com/noobclear/nuclear-bot/app/request"
	"github.com/sirupsen/logrus"
	"net"
	"net/textproto"
)

type Bot interface {
	Start()
}

type NuclearBot struct {
	Config  config.BotConfig
	Handler handlers.Handler
}

func (b *NuclearBot) Start() {
	conn := b.getConnection()
	defer conn.Close()

	r := textproto.NewReader(bufio.NewReader(conn))
	ctx := request.Context{
		BotUsername:    b.Config.BotUsername,
		TargetChannel:  b.Config.TargetChannel,
		Clients:        clients.NewClients(b.Config),
		CommandFactory: commands.NewCommandFactory(),
	}

	w := msgs.NewMessageWriter(conn, b.Config.RateLimit)
	b.start(&ctx, r, w)
}

func (b *NuclearBot) getConnection() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")

	if err != nil {
		panic(err)
	}

	// token, username, and channel
	conn.Write([]byte("PASS " + b.Config.TwitchOAuthToken + "\r\n"))
	conn.Write([]byte("NICK " + b.Config.BotUsername + "\r\n"))
	conn.Write([]byte("JOIN " + b.Config.TargetChannel + "\r\n"))

	logrus.Infof("%s joined %s", b.Config.BotUsername, b.Config.TargetChannel)
	return conn
}

func (b *NuclearBot) start(ctx *request.Context, r *textproto.Reader, w msgs.Writer) {
	var lineCount uint32

	for {
		line, err := r.ReadLine()
		if err != nil {
			panic(err)
		}

		lineCount++
		logrus.Infof("%d< [%s]", lineCount, line)

		msg, err := msgs.NewMessage(line)
		if err != nil {
			logrus.Warn(err.Error())
		} else {
			b.Handler.Handle(ctx, w, msg)
		}
	}
}
