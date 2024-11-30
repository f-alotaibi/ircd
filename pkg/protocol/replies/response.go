package replies

import "fmt"

var (
	RPL_NONE     = ""
	RPL_USERHOST = ""
	RPL_ISON     = ""
	RPL_AWAY     = func(nick string, awayMessage string) []byte {
		return []byte(fmt.Sprintf("301 %s :%s", nick, awayMessage))
	}
	RPL_UNAWAY        = []byte("305 :You are no longer marked as being away")
	RPL_NOWAWAY       = []byte("306 :You have been marked as being away")
	RPL_WHOISUSER     = ""
	RPL_WHOISSERVER   = ""
	RPL_WHOISOPERATOR = ""
	RPL_WHOISIDLE     = ""
	RPL_ENDOFWHOIS    = ""
	RPL_WHOISCHANNELS = ""
	RPL_WHOWASUSER    = ""
	RPL_ENDOFWHOWAS   = ""
	RPL_LISTSTART     = ""
	RPL_LIST          = ""
	RPL_LISTEND       = ""
	RPL_CHANNELMODEIS = ""
	RPL_NOTOPIC       = func(channel string) []byte { return []byte(fmt.Sprintf("331 %s :No topic is set", channel)) }
	RPL_TOPIC         = func(channel string, topic string) []byte { return []byte(fmt.Sprintf("332 %s :%s", channel, topic)) }
	RPL_INVITING      = ""
	RPL_SUMMONING     = ""
	RPL_VERSION       = ""
	RPL_WHOREPLY      = ""
	RPL_ENDOFWHO      = ""
	RPL_NAMREPLY      = ""
	RPL_ENDOFNAMES    = ""
	RPL_LINKS         = ""
	RPL_ENDOFLINKS    = ""
	RPL_BANLIST       = ""
	RPL_ENDOFBANLIST  = ""
)
