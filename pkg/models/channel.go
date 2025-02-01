package models

import (
	"fmt"
	"slices"
	"sync"
)

type Channel struct {
	Name    string
	Topic   string
	syncRW  sync.RWMutex
	Clients []*Client
}

func (ch *Channel) Add(client *Client) error {
	ch.syncRW.RLock()
	ok := slices.Contains(ch.Clients, client)
	ch.syncRW.RUnlock()
	if ok {
		return fmt.Errorf("user already exists")
	}
	ch.syncRW.Lock()
	ch.Clients = append(ch.Clients, client)
	ch.syncRW.Unlock()
	return nil
}

func (ch *Channel) Remove(client *Client) error {
	ch.syncRW.RLock()
	i := slices.Index(ch.Clients, client)
	ch.syncRW.RUnlock()
	if i == -1 {
		return fmt.Errorf("user doesn't exist")
	}
	ch.syncRW.Lock()
	ch.Clients = append(ch.Clients[:i], ch.Clients[i+1:]...)
	ch.syncRW.Unlock()
	return nil
}

func (ch *Channel) GetClients() []*Client {
	ch.syncRW.RLock()
	clients := make([]*Client, len(ch.Clients))
	copy(clients, ch.Clients)
	ch.syncRW.RUnlock()
	return clients
}
