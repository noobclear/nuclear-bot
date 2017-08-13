package messages

import "net"

type Context struct {
	Connection net.Conn
	BotUsername string
	TargetChannel string
}