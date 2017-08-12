package app

import "strings"

type Responder interface {
	Respond(msg string, ctx Context)
}

type MessageResponder struct {}

func (mh *MessageResponder) Respond(msg string, ctx Context) {
	// TODO: parse and handle message

	// i.e. ":username@username.tmi.twitch.tv PRIVMSG #channel :chatmessage"
	msgParts := strings.Split(msg, " ")

	if msgParts[0] == "PING" {
		ctx.Connection.Write([]byte("PONG " + msgParts[1] + CRLF))
		return
	}

	if msgParts[1] == "PRIVMSG" {
		ctx.Connection.Write([]byte("PRIVMSG " + msgParts[2] + " " + msgParts[3] + CRLF))
	}
}

func NewMessageResponder() Responder {
	return &MessageResponder{}
}