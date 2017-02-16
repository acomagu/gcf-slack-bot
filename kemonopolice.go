package main

import (
	"fmt"
	"github.com/acomagu/chatroom-go/chatroom"
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

func kemonoPoliceTopic(room chatroom.Room) chatroom.DidTalk {
	r := waitReceived(room)
	if r.channelName != "kemono" || isLegal(r.text) {
		return false
	}
	_, _, err := api.DeleteMessage(r.channelID, r.timestamp)
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
