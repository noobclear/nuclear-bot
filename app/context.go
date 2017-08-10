package app

import "net"

type Context struct {
	Connection net.Conn
}