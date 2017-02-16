package main

import (
	"fmt"
	"github.com/acomagu/chatroom-go/chatroom"
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

func kemonoReactionTopic(room chatroom.Room) chatroom.DidTalk {
	r := waitReceived(room)
	if r.channelName == "kemono" || !doesIncludeKemonoWords(r.text) {
		return false
	}
	theItem := slack.ItemRef{
		Channel:   r.channelID,
		Timestamp: r.timestamp,
	}
	err := api.AddReaction("kemono_friends", theItem)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func doesIncludeKemonoWords(msg string) bool {
	for _, kemonoWord := range kemonoWords {
		if kemonoWord.MatchString(msg) {
			return true
		}
	}
	return false
}
