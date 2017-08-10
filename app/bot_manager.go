package app

import (
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

func NewBotManager(c *Config) Manager {
	var bots []Starter
	for _, bc := range c.BotConfigs {
		bot := Bot{bc, NewMessageResponder()}
		bots = append(bots, &bot)
	}
	return &BotManager{bots}
}
