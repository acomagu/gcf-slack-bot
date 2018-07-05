package nanhankanmu

import (
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/topicutil"
	"regexp"
)

var trigger = regexp.MustCompile(`3[^\w]*3[^\w]*4`)

// Client keeps slack client to delete other's message.
type Client struct{}

// Talk is main Topic
func (Client) Talk(room chatroom.Room) chatroom.DidTalk {
	r := topicutil.WaitReceived(room)
	if !trigger.MatchString(r.Text) {
		return false
	}

	room.Send("なんでや! 阪神関係ないやろ!")
	return true
}
