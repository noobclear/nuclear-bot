package app

import "sync"

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

func NewBotManager(c *Config) *BotManager {
	var bots []Bot
	for _, bc := range c.BotConfigs {
		bot := Bot{bc}
		bots = append(bots, bot)
	}
	return &BotManager{bots}
}
