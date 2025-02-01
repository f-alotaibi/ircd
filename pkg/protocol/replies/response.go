package replies

import (
	"fmt"
	"strings"
)

var (
	RPL_LUSERCLIENT = func(users int, invisible int, servers int) []byte {
		return []byte(fmt.Sprintf("251 :There are %d users and %d invisible on %d servers", users, invisible, servers))
	}
	RPL_NONE = []byte("300")
	RPL_AWAY = func(nick string, awayMessage string) []byte {
		return []byte(fmt.Sprintf("301 %s :%s", nick, awayMessage))
	}
	RPL_UNAWAY  = []byte("305 :You are no longer marked as being away")
	RPL_NOWAWAY = []byte("306 :You have been marked as being away")
	RPL_TOPIC   = func(channel string, topic string) []byte {
		return []byte(fmt.Sprintf("332 %s :%s", channel, topic))
	}
	RPL_NAMREPLY = func(channel string, users []string) []byte {
		return []byte(fmt.Sprintf("353 %s :%s", channel, strings.Join(users, " ")))
	}
	RPL_MOTD          = func(text string) []byte { return []byte(fmt.Sprintf("372 :- %s", text)) }
	RPL_MOTDSTART     = func(server string) []byte { return []byte(fmt.Sprintf("375 :- %s Message of the day - ", server)) }
	RPL_NOTOPIC       = func(channel string) []byte { return []byte(fmt.Sprintf("331 %s :No topic is set", channel)) }
	RPL_ENDOFMOTD     = []byte("376 :End of /MOTD command")
	RPL_USERHOST      = ""
	RPL_ISON          = ""
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
	RPL_INVITING      = ""
	RPL_SUMMONING     = ""
	RPL_VERSION       = ""
	RPL_WHOREPLY      = ""
	RPL_ENDOFWHO      = ""
	RPL_ENDOFNAMES    = ""
	RPL_LINKS         = ""
	RPL_ENDOFLINKS    = ""
	RPL_BANLIST       = ""
	RPL_ENDOFBANLIST  = ""
)
