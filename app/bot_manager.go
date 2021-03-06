package app

import (
	"github.com/noobclear/nuclear-bot/app/config"
	h "github.com/noobclear/nuclear-bot/app/handlers"
	"github.com/noobclear/nuclear-bot/app/msgs"
	"github.com/noobclear/nuclear-bot/app/request"
	"sync"
)

type Manager interface {
	StartAll()
}

type BotManager struct {
	Bots []Bot
}

func (bm *BotManager) StartAll() {
	var wg sync.WaitGroup
	wg.Add(len(bm.Bots))
	for _, b := range bm.Bots {
		go b.Start()
	}
	wg.Wait()
}

func NewBotManager(c *config.Config) Manager {
	var bots []Bot

	for _, bc := range c.BotConfigs {
		bot := NuclearBot{
			Config:  bc,
			Handler: adapt(h.NewPingHandler, h.NewIgnoreBotsHandler, h.NewCommandHandler, h.NewNLPHandler),
		}
		bots = append(bots, &bot)
	}
	return &BotManager{bots}
}

func adapt(handlers ...func(h.Handler) h.Handler) h.Handler {
	var curr h.Handler
	for i := len(handlers) - 1; i >= 0; i-- {
		if curr == nil {
			end := h.HandlerFunc(func(*request.Context, msgs.Writer, msgs.Message) {})
			curr = handlers[i](end)
		} else {
			curr = handlers[i](curr)
		}
	}
	return curr
}
