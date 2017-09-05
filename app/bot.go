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
	"time"
	"github.com/beefsack/go-rate"
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

	clients := clients.NewClients(b.Config)
	r := textproto.NewReader(bufio.NewReader(conn))
	ctx := request.Context{
		BotUsername:    b.Config.BotUsername,
		TargetChannel:  b.Config.TargetChannel,
		Clients:        clients,
		CommandFactory: commands.NewCommandFactory(clients),
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
	go alertNewFollowers(ctx)
	receiveIRCChatMessages(b.Handler, ctx, r, w)
}

func receiveIRCChatMessages(h handlers.Handler, ctx *request.Context, r *textproto.Reader, w msgs.Writer) {
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
			h.Handle(ctx, w, msg)
		}
	}
}

func alertNewFollowers(ctx *request.Context) {
	// Poll for new followers
	r := rate.New(1, time.Second * 30)

	// Get user ID from channel
	user, err := ctx.Clients.Twitch.GetUserByLogin(ctx.TargetChannel[1:])
	if err != nil || user == nil || user.ID == "" {
		panic(err)
	}

	// Doesn't scale with millions of followers but definitely enough for our needs
	followers := make(map[string]struct{})

	// Build initial list of latest 100 followers
	resp, err := ctx.Clients.Twitch.GetChannelFollows(user.ID)
	if err != nil {
		logrus.Warnf("Failed to get initial channel follows for %s", user.ID)
		panic(err)
	}
	for _, f := range resp.Follows {
		logrus.Infof("Adding %+v to initial followers list", f.User)
		followers[f.User.ID] = struct{}{}
	}
	logrus.Infof("Added %v followers", len(followers))
	r.Wait()

	for {
		r.Wait()
		resp, err := ctx.Clients.Twitch.GetChannelFollows(user.ID)
		if err != nil {
			logrus.Warnf("Failed to get channel follows for %s, continuing...", user.ID)
			continue
		}

		for _, f := range resp.Follows {
			_, ok := followers[f.User.ID]
			if !ok {
				// Trigger new follower alert and add to list
				logrus.Infof("New follower alert: %+v", f.User)
				// TODO: rate limit new follower alert with channel and call streamlabs
				ctx.Clients.Lights.TriggerPulse()
				followers[f.User.ID] = struct{}{}
			}
		}
	}
}