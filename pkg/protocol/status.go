package protocol

type Status uint8

const (
	NOT_REGISTERED Status = iota
	NICK_OK
	USER_OK
	REGISTERED
)
