package channelhost

import "irc/pkg/models"

type ChannelHost interface {
	Channels() map[string]*models.Channel
}
