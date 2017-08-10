package app

import (
	"bufio"
	"net"
	"net/textproto"
)

const (
	CRLF = "\r\n"
)

type Starter interface {
	Start()
}

type Bot struct {
	BotConfig
	Responder
}

func (b *Bot) Start() {
	conn := b.GetConnection()
	defer conn.Close()

	tp := textproto.NewReader(bufio.NewReader(conn))
	ctx := Context{conn}

	for {
		msg, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		b.Responder.Respond(msg, ctx)
	}
}

func (b *Bot) GetConnection() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")

	if err != nil {
		panic(err)
	}

	// token, username, and channel
	conn.Write([]byte("PASS " + b.BotConfig.TwitchOAuthToken + CRLF))
	conn.Write([]byte("NICK " + b.BotConfig.BotUsername + CRLF))
	conn.Write([]byte("JOIN " + b.BotConfig.TargetChannel + CRLF))
	return conn
}
