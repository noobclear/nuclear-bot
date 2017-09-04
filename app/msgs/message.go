package msgs

import (
	"errors"
	"fmt"
	"strings"
)

type Message interface {
	FormatResponse() string
}

func NewMessage(msg string) (Message, error) {
	// msg looks like :nuclear!nuclear@nuclear.tmi.twitch.tv PRIVMSG #nuclear :some message here
	msg = strings.TrimSpace(msg)
	msgParts := strings.Split(msg, " ")

	if len(msgParts) > 0 && msgParts[0] == "PING" {
		return &PingMessage{Host: msgParts[1]}, nil
	} else if len(msgParts) > 3 && msgParts[1] == "PRIVMSG" {
		return &PrivMessage{
			// i.e. Extract "nuclear" from ":nuclear!nuclear@nuclear.tmi.twitch.tv"
			FromUser: msgParts[0][1:strings.Index(msgParts[0], "!")],
			// i.e. Extract "#nuclear"
			ToChannel: msgParts[2],
			Text:      strings.SplitN(msg, ":", 3)[2],
		}, nil
	}

	return nil, errors.New(fmt.Sprintf("Unable to parse message: [%s]", msg))
}
