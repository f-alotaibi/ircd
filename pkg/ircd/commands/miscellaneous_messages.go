package commands

import "irc/pkg/models"

type PingCommand struct{}

func (j PingCommand) Handle(client *models.Client, args []string) {

}
