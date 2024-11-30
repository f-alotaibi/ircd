package models

import "net"

type Client struct {
	Conn net.Conn
	User User
}

func (c *Client) Respond(bytes []byte) {
	c.Conn.Write(bytes)
}
