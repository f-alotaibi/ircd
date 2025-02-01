package commands

import (
	"irc/pkg/ircd/channelhost"
	"irc/pkg/ircd/handler"
)

var DefaultCommands = func(ch channelhost.ChannelHost) map[string]handler.CommandHandler {
	return map[string]handler.CommandHandler{
		"PASS":   PassCommand{},
		"NICK":   NickCommand{},
		"USER":   UserCommand{},
		"SERVER": ServerCommand{},
		"OPER":   OperCommand{},
		"SQUIT":  ServerQuitCommand{},
		"JOIN":   JoinCommand{ChannelHost: ch},
		"PART":   PartCommand{ChannelHost: ch},
	}
}
