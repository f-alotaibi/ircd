package commands

import (
	"irc/pkg/ircd/channelhost"
	"irc/pkg/models"
	"irc/pkg/protocol/replies"
	"strings"
)

type JoinCommand struct {
	ChannelHost channelhost.ChannelHost
}

func (j JoinCommand) Handle(client *models.Client, args []string) {
	if len(args) < 1 {
		client.Respond([]byte(replies.ERR_NEEDMOREPARAMS("JOIN")))
		return
	}
	splitChannels := strings.Split(args[0], ",")
	for _, channelName := range splitChannels {
		// TODO: KEY
		channel, ok := j.ChannelHost.Channels()[channelName]
		if !ok {
			continue
		}
		err := channel.Add(client)
		if err != nil {
			continue
		}
		if channel.Topic == "" {
			client.Respond(replies.RPL_NOTOPIC(channel.Name))
		} else {
			client.Respond(replies.RPL_TOPIC(channel.Name, channel.Topic))
		}
	}
}

type PartCommand struct {
	ChannelHost channelhost.ChannelHost
}

func (j PartCommand) Handle(client *models.Client, args []string) {
	if len(args) < 1 {
		client.Respond([]byte(replies.ERR_NEEDMOREPARAMS("PART")))
		return
	}
	splitChannels := strings.Split(args[0], ",")
	for _, channelName := range splitChannels {
		// TODO: KEY
		channel, ok := j.ChannelHost.Channels()[channelName]
		if !ok {
			continue
		}
		err := channel.Remove(client)
		if err != nil {
			continue
		}
	}
}
