package main

import (
	"regexp"
	"github.com/acomagu/chatroom-go/chatroom"
)

var topics = []chatroom.Topic{kemonoPoliceTopic, kemonoReactionTopic, suggestRestaurant}

func waitReceived(room chatroom.Room) received {
	for {
		msg := room.WaitMsg()
		if r, ok := msg.(received); ok {
			return r
		}
	}
}

func matchAny(regexps []*regexp.Regexp, str string) bool {
	for _, exp := range regexps {
		if exp.MatchString(str) {
			return true
		}
	}
	return false
}
