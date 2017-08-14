package messages

import (
	"strings"
	"github.com/sirupsen/logrus"
)

type Message interface {}

type PingMessage struct {
	Host string
}

type PrivMessage struct {
	FromUser string
	ToChannel string
	Text string
}

func NewMessage(msg string) Message {
	// msg looks like :nuclear!nuclear@nuclear.tmi.twitch.tv PRIVMSG #nuclear :some message here
	msgParts := strings.Split(msg, " ")

	if len(msgParts) > 0 && msgParts[0] == "PING" {
		return &PingMessage{Host: msgParts[1]}
	} else if len(msgParts) > 3 && msgParts[1] == "PRIVMSG" {
		return &PrivMessage{
			// i.e. Extract "nuclear" from ":nuclear!nuclear@nuclear.tmi.twitch.tv"
			FromUser: msgParts[0][1:strings.Index(msgParts[0], "!")],
			// i.e. Extract "#nuclear"
			ToChannel: msgParts[2],
			Text: strings.SplitN(msg, ":", 3)[2],
		}
	}

	logrus.Warnf("Unable to parse message: %s", msg)
	return nil
}