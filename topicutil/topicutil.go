package topicutil

import (
	"regexp"
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/slackcr"
)

// WaitReceived is wrap of chatroom.Room.WaitMsg. Type annotacing for slackcr.Received type and returns it.
func WaitReceived(room chatroom.Room) slackcr.Received {
	for {
		msg := room.WaitMsg()
		if r, ok := msg.(slackcr.Received); ok {
			return r
		}
	}
}

// MatchAny trys all of reglar expression in slice to str, and if any matched, return true.
func MatchAny(regexps []*regexp.Regexp, str string) bool {
	for _, exp := range regexps {
		if exp.MatchString(str) {
			return true
		}
	}
	return false
}
