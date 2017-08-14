package app

import (
	"sync"
	"github.com/noobclear/nuclear-bot/app/config"
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
	for _, bc := range c.BotConfigs {
		bot := Bot{bc, NewMessageRouter()}
		bots = append(bots, &bot)
	}
	return &BotManager{bots}
}
