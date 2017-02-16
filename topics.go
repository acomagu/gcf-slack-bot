package main

import (
	"github.com/acomagu/chatroom-go/chatroom"
)

var topics = []chatroom.Topic{kemonoPoliceTopic, kemonoReactionTopic}

func waitReceived(room chatroom.Room) received {
	for {
		msg := room.WaitMsg()
		if r, ok := msg.(received); ok {
			return r
		}
	}
}
