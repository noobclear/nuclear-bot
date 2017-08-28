package msgs

import (
	"github.com/beefsack/go-rate"
	"github.com/noobclear/nuclear-bot/app/util"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type Writer interface {
	Write(msg string)
}

type MessageWriter struct {
	ResponseChannel chan<- string
}

func (w *MessageWriter) Write(msg string) {
	if msg != "" {
		w.ResponseChannel <- msg
	}
}

func NewMessageWriter(c net.Conn, rateLimit int) Writer {
	// Create a rate limited channel to process bot responses
	limiter := rate.New(rateLimit, 30*time.Second)
	messageQueue := make(chan string, rateLimit)
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
			c.Write([]byte(s + util.CRLF))
		}
	}(messageQueue)

	return &MessageWriter{
		ResponseChannel: messageQueue,
	}
}
