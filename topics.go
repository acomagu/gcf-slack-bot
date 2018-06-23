package main

import (
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/slackcr"
	"github.com/acomagu/gcf-slack-bot/restaurants"
	"github.com/acomagu/gcf-slack-bot/kmnreact"
	"github.com/acomagu/gcf-slack-bot/nanhankanmu"
	"github.com/acomagu/gcf-slack-bot/twada"
)

func topics(clients slackcr.SlackClients) []chatroom.Topic {
	rests := restaurants.New(clients.Friends)
	react := kmnreact.New(clients.God)
	nnkm := nanhankanmu.Client{}
	twada := twada.Client{}
	return []chatroom.Topic{rests.Talk, react.Talk, nnkm.Talk, twada.Talk}
}
