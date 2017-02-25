package kmnreact

import (
	"fmt"
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/topicutil"
	"github.com/nlopes/slack"
	"regexp"
)

var kemonoWords = []*regexp.Regexp{
	regexp.MustCompile(`すっ?ごーい`),
	regexp.MustCompile(`たー?のしー`),
	regexp.MustCompile(`わーい`),
	regexp.MustCompile(`へーきへーき`),
	regexp.MustCompile(`フレンズ`),
	regexp.MustCompile(`うわー`),
	regexp.MustCompile(`なにこれなにこれ`),
	regexp.MustCompile(`おもしろーい`),
}

// Client keeps slack client to delete other's message.
type Client struct {
	bot *slack.Client
}

// New creates new Client.
func New(bot *slack.Client) Client {
	return Client{
		bot: bot,
	}
}

// Talk is main Topic.
func (client Client) Talk(room chatroom.Room) chatroom.DidTalk {
	r := topicutil.WaitReceived(room)
	if r.ChannelName == "kemono" || !doesIncludeKemonoWords(r.Text) {
		return false
	}
	theItem := slack.ItemRef{
		Channel:   r.ChannelID,
		Timestamp: r.Timestamp,
	}
	err := client.bot.AddReaction("kemono_friends", theItem)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func doesIncludeKemonoWords(msg string) bool {
	return topicutil.MatchAny(kemonoWords, msg)
}
