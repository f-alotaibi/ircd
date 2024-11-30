package handler

import "irc/pkg/models"

type Command interface {
	Handle(*models.Client, []string)
}
