package commands

import (
	"irc/pkg/ircd/channelhost"
	"irc/pkg/ircd/handler"
)

var DefaultCommands = func(ch channelhost.ChannelHost) map[string]handler.Command {
	return map[string]handler.Command{
		"PASS":   PassCommand{},
		"NICK":   NickCommand{},
		"USER":   UserCommand{},
		"SERVER": ServerCommand{},
		"OPER":   OperCommand{},
		"SQUIT":  ServerQuitCommand{},
		"JOIN":   JoinCommand{ChannelHost: ch},
	}
}
