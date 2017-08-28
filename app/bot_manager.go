package app

import (
	"github.com/noobclear/nuclear-bot/app/config"
	h "github.com/noobclear/nuclear-bot/app/handlers"
	"github.com/noobclear/nuclear-bot/app/msgs"
	"sync"
)

type Manager interface {
	StartAll()
}

type BotManager struct {
	Bots []Starter
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
	var bots []Starter
	end := h.HandlerFunc(func(*msgs.Context, msgs.Writer, msgs.Message) {})

	for _, bc := range c.BotConfigs {
		bot := Bot{bc, h.NewPingHandler(h.NewIgnoreSelfHandler(h.NewNLPHandler(end)))}
		bots = append(bots, &bot)
	}
	return &BotManager{bots}
}
