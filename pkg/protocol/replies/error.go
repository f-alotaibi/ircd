package replies

import "fmt"

// https://datatracker.ietf.org/doc/html/rfc1459#section-6.1
var (
	ERR_NOSUCHNICK       = func(nick string) []byte { return []byte(fmt.Sprintf("401 %s :No such nick/channel", nick)) }
	ERR_NOSUCHSERVER     = func(server string) []byte { return []byte(fmt.Sprintf("402 %s :No such server", server)) }
	ERR_NOSUCHCHANNEL    = func(channel string) []byte { return []byte(fmt.Sprintf("403 %s :No such channel", channel)) }
	ERR_CANNOTSENDTOCHAN = func(channel string) []byte { return []byte(fmt.Sprintf("404 %s :Cannot send to channel", channel)) }
	ERR_TOOMANYCHANNELS  = func(channel string) []byte {
		return []byte(fmt.Sprintf("405 %s :You have joined too many channels", channel))
	}
	ERR_WASNOSUCHNICK  = func(nick string) []byte { return []byte(fmt.Sprintf("406 %s :There was no such nickname", nick)) }
	ERR_TOOMANYTARGETS = func(target string) []byte {
		return []byte(fmt.Sprintf("407 %s :Duplicate recipients. No message delivered", target))
	}
	ERR_NONICKNAMEGIVEN  = []byte("431 :No nickname given")
	ERR_ERRONEUSNICKNAME = func(nick string) []byte { return []byte(fmt.Sprintf("432 %s :Erroneus nickname", nick)) }
	ERR_NICKNAMEINUSE    = func(nick string) []byte { return []byte(fmt.Sprintf("433 %s :Nickname is already in use", nick)) }
	ERR_NICKCOLLISION    = func(nick string) []byte { return []byte(fmt.Sprintf("436 %s :Nickname collision KILL", nick)) }
	ERR_NOTONCHANNEL     = func(channel string) []byte { return []byte(fmt.Sprintf("442 %s :You're not on that channel", channel)) }
	ERR_NEEDMOREPARAMS   = func(cmd string) []byte { return []byte(fmt.Sprintf("461 %s :Not enough parameters", cmd)) }
	ERR_ALREADYREGISTRED = []byte("462 :You may not reregister")
	ERR_PASSWDMISMATCH   = []byte("464 :Password incorrect")
	ERR_KEYSET           = func(channel string) []byte { return []byte(fmt.Sprintf("467 %s :Channel key already set", channel)) }
	ERR_CHANNELISFULL    = func(channel string) []byte { return []byte(fmt.Sprintf("471 %s :Cannot join channel (+l)", channel)) }
	ERR_UNKNOWNMODE      = func(char string) []byte { return []byte(fmt.Sprintf("472 %s :is unknown mode char to me", char)) }
	ERR_INVITEONLYCHAN   = func(channel string) []byte { return []byte(fmt.Sprintf("473 %s :Cannot join channel (+i)", channel)) }
	ERR_BANNEDFROMCHAN   = func(channel string) []byte { return []byte(fmt.Sprintf("474 %s :Cannot join channel (+b)", channel)) }
	ERR_BADCHANNELKEY    = func(channel string) []byte { return []byte(fmt.Sprintf("475 %s :Cannot join channel (+k)", channel)) }
	ERR_NOPRIVILEGES     = []byte("481 :Permission Denied- You're not an IRC operator")
	ERR_NOOPERHOST       = []byte("491 :No O-lines for your host")
	ERR_BADCHANMASK      = []byte("476")
	ERR_CHANOPRIVSNEEDED = func(channel string) []byte {
		return []byte(fmt.Sprintf("482 %s :You're not channel operator", channel))
	}
	ERR_USERSDONTMATCH   = []byte("502 :Cant change mode for other users")
	ERR_UMODEUNKNOWNFLAG = []byte("501 :Unknown MODE flag")
)
