package config

import (
	"encoding/json"
	"io/ioutil"
)

const (
	ConfigFile = "config.json"
)

// Example config.json file:
//
//{
//  "bots": [
//    {
//      "twitch_oauth_token": "oauth:<OAUTH_TOKEN>",
//      "bot_username": "newbclear",
//      "target_channel": "#nuclear",
//      "rate_limit": 100 // 100 msgs per 30s
//    }
//  ]
//}
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
