package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/textproto"
	"strings"
)

const (
	CRLF       = "\r\n"
	ConfigFile = "config.json"
)

type Config struct {
	TwitchOAuthToken string `json:"twitch_oauth_token"`
	BotUsername      string `json:"bot_username"`
	TargetChannel    string `json:"target_channel"`
}

func getConfig() *Config {
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

func main() {
	conf := getConfig()

	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		panic(err)
	}

	// token, username, and channel
	conn.Write([]byte("PASS " + conf.TwitchOAuthToken + CRLF))
	conn.Write([]byte("NICK " + conf.BotUsername + CRLF))
	conn.Write([]byte("JOIN " + conf.TargetChannel + CRLF))
	defer conn.Close()

	tp := textproto.NewReader(bufio.NewReader(conn))

	for {
		msg, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}

		// i.e. ":username@username.tmi.twitch.tv PRIVMSG #channel :chatmessage"
		msgParts := strings.Split(msg, " ")

		if msgParts[0] == "PING" {
			conn.Write([]byte("PONG " + msgParts[1]))
			continue
		}

		if msgParts[1] == "PRIVMSG" {
			conn.Write([]byte("PRIVMSG " + msgParts[2] + " " + msgParts[3] + "\r\n"))
		}
	}
}
