package app

import (
	"net"
	"net/textproto"
	"bufio"
	"strings"
)

const (
	CRLF       = "\r\n"
)

type Bot struct {
	BotConfig
}

func (b *Bot) Start() {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")

	if err != nil {
		panic(err)
	}

	// token, username, and channel
	conn.Write([]byte("PASS " + b.BotConfig.TwitchOAuthToken + CRLF))
	conn.Write([]byte("NICK " + b.BotConfig.BotUsername + CRLF))
	conn.Write([]byte("JOIN " + b.BotConfig.TargetChannel + CRLF))
	defer conn.Close()

	tp := textproto.NewReader(bufio.NewReader(conn))

	for {
		msg, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}

		// i.e. ":username@username.tmi.twitch.tv PRIVMSG #channel :chatmessage"
		msgParts := strings.Split(msg, " ")

		if msgParts[0] == "PING" {
			conn.Write([]byte("PONG " + msgParts[1]))
			continue
		}

		if msgParts[1] == "PRIVMSG" {
			conn.Write([]byte("PRIVMSG " + msgParts[2] + " " + msgParts[3] + "\r\n"))
		}
	}
}