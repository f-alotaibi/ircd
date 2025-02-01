package models

import (
	"irc/pkg/protocol"
	"net"
)

type Client struct {
	Conn   net.Conn
	User   User
	Status protocol.Status
}

func (c *Client) Respond(bytes []byte) {
	response := append([]byte(":127.0.0.1 "), bytes...)
	response = append(response, 0x0d, 0x0a)
	c.Conn.Write(response)
}
