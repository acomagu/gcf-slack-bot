package main

import (
	"fmt"
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/topicutil"
	"github.com/nlopes/slack"
	"regexp"
)

var laws = []*regexp.Regexp{
	regexp.MustCompile(`すっ?ごーい`),
	regexp.MustCompile(`たー?のしー`),
	regexp.MustCompile(`わーい`),
	regexp.MustCompile(`へーきへーき`),
	regexp.MustCompile(`[^!！。、]+ね`),
	regexp.MustCompile(`フレンズによって(得意|とくい)なこと(違|ちが)うから`),
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

// Talk is main Topic
func (client Client) Talk(room chatroom.Room) chatroom.DidTalk {
	r := topicutil.WaitReceived(room)
	if r.ChannelName != "kemono" || isLegal(r.Text) {
		return false
	}
	_, _, err := client.bot.DeleteMessage(r.ChannelID, r.Timestamp)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func isLegal(msg string) bool {
	removes := regexp.MustCompile(`(\s|\t|\n|!|！)`)
	msg = removes.ReplaceAllString(msg, "")

	for _, law := range laws {
		msg = law.ReplaceAllString(msg, "")
	}
	return msg == ""
}
