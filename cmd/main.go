package main

import (
	"irc/pkg/ircd"
	"irc/pkg/models"
)

func main() {
	ircd := ircd.New("0.0.0.0:6667")
	ircd.ChannelMap = map[string]*models.Channel{
		"test": {
			Name: "test",
		},
	}
	err := ircd.Listen()
	if err != nil {
		panic(err)
	}
}
