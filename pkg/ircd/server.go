package ircd

import (
	"irc/pkg/ircd/commands"
	"irc/pkg/ircd/handler"
	"irc/pkg/models"
	"log"
	"net"
	"strings"
)

type Server struct {
	Address    string
	Logger     *log.Logger
	Handler    map[string]handler.Command
	ChannelMap map[string]*models.Channel
}

func New(Address string) *Server {
	server := &Server{
		Address: Address,
		Logger:  log.Default(),
	}
	server.Handler = commands.DefaultCommands(server)
	return server
}

func (l Server) Listen() error {
	ln, err := net.Listen("tcp", l.Address)
	if err != nil {
		return err
	}
	defer ln.Close()
	l.Logger.Printf("Server is listening at %s\n", l.Address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			l.Logger.Println("Error handling")
			continue
		}
		go l.handleConnection(conn)
	}
}

func (l Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	client := models.Client{
		Conn: conn,
	}
	buffer := make([]byte, 512)
	for {
		// Read packets
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		plainText := string(buffer[:n-2])
		requests := strings.Split(plainText, "\n")
		for _, request := range requests {
			requestSplit := strings.Split(request, " ")

			cmdName := requestSplit[0]
			var args []string
			if len(requestSplit) > 1 {
				args = requestSplit[1:]
			}
			cmd, ok := l.Handler[cmdName]
			if !ok {
				l.Logger.Printf("Handler for %s not found, args are %v, buffer is %s\n", cmdName, args, plainText)
				continue
			}
			l.Logger.Printf("Handling %s with args %v\n", cmdName, args)
			cmd.Handle(&client, args)
		}
	}
}

func (l Server) Channels() map[string]*models.Channel {
	return l.ChannelMap
}
