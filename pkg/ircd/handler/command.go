package handler

import "irc/pkg/models"

type CommandHandler interface {
	Handle(*models.Client, []string)
}
