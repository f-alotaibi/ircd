package commands

import (
	"irc/pkg/models"
	"irc/pkg/protocol/replies"
)

type PassCommand struct {
}

func (j PassCommand) Handle(client *models.Client, args []string) {
	if len(args) < 1 {
		client.Respond([]byte(replies.ERR_NEEDMOREPARAMS("PASS")))
		return
	}
	client.User.Password = args[0]
}

type NickCommand struct {
}

func (j NickCommand) Handle(client *models.Client, args []string) {
	if len(args) < 1 {
		client.Respond([]byte(replies.ERR_NONICKNAMEGIVEN))
		return
	}
	client.User.Nick = args[0]
}

type UserCommand struct {
}

func (j UserCommand) Handle(client *models.Client, args []string) {
	if len(args) < 4 {
		client.Respond([]byte(replies.ERR_NEEDMOREPARAMS("USER")))
		return
	}
	client.User.Name = args[0]
	client.User.Realname = args[3]
}

type ServerCommand struct {
}

func (j ServerCommand) Handle(client *models.Client, args []string) {
	// TODO
}

type OperCommand struct {
}

func (j OperCommand) Handle(client *models.Client, args []string) {
	// TODO
}

type QuitCommand struct {
}

func (j QuitCommand) Handle(client *models.Client, args []string) {
	// TODO
}

type ServerQuitCommand struct {
}

func (j ServerQuitCommand) Handle(client *models.Client, args []string) {
	// TODO
}
