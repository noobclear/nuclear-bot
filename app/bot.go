package app

import (
	"bufio"
	"github.com/beefsack/go-rate"
	"github.com/noobclear/nuclear-bot/app/config"
	"github.com/noobclear/nuclear-bot/app/messages"
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/sirupsen/logrus"
	"net"
	"net/textproto"
	"time"
)

type Starter interface {
	Start()
}

type Bot struct {
	Config config.BotConfig
	Router Router
}

func (b *Bot) Start() {
	conn := b.getConnection()
	defer conn.Close()

	r := textproto.NewReader(bufio.NewReader(conn))
	ctx := messages.Context{
		Connection:    conn,
		BotUsername:   b.Config.BotUsername,
		TargetChannel: b.Config.TargetChannel,
	}

	// TODO: Refactor out this rate limited channel
	// Create a rate limited channel to process bot responses
	limiter := rate.New(b.Config.RateLimit, 30*time.Second)
	messageQueue := make(chan string, b.Config.RateLimit)
	var count int

	go func(q <-chan string) {
		for s := range q {
			count++
			ok, remaining := limiter.Try()
			// Throw away response if user has to wait too long
			if !ok {
				logrus.Warnf("Rate limited skip: [%s], remaining: %v", s, remaining)
				continue
			}
			limiter.Wait()
			logrus.Infof("%d> [%s]", count, s)
			conn.Write([]byte(s + util.CRLF))
		}
	}(messageQueue)

	b.listen(&ctx, r, messageQueue)
}

func (b *Bot) getConnection() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")

	if err != nil {
		panic(err)
	}

	// token, username, and channel
	conn.Write([]byte("PASS " + b.Config.TwitchOAuthToken + util.CRLF))
	conn.Write([]byte("NICK " + b.Config.BotUsername + util.CRLF))
	conn.Write([]byte("JOIN " + b.Config.TargetChannel + util.CRLF))

	logrus.Infof("%s joined %s", b.Config.BotUsername, b.Config.TargetChannel)
	return conn
}

func (b *Bot) listen(ctx *messages.Context, r *textproto.Reader, q chan<- string) {
	for {
		msg, err := r.ReadLine()
		if err != nil {
			panic(err)
		}

		logrus.Infof("< %s", msg)
		b.Router.Route(ctx, msg, q)
	}
}
