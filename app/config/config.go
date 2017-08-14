package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	ConfigFile = "config.json"
)

type Config struct {
	BotConfigs []BotConfig `json:"bots"`
}

func GetConfig() *Config {
	f, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}

	c := Config{}
	err = json.Unmarshal(f, &c)
	if err != nil {
		panic(err)
	}

	return &c
}
