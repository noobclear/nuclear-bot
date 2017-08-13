package messages

import "strings"

type Message struct {
	Original string
	MsgParts []string
}

func NewMessage(msg string) Message {
	return Message{
		Original: msg,
		// i.e. ":username@username.tmi.twitch.tv PRIVMSG #channel :chatmessage"
		MsgParts: strings.Split(msg, " "),
	}
}