package main

import (
	"os"

	"github.com/acomagu/gcf-slack-bot/slackcr"
)

var port = os.Getenv("PORT")
var botAPIToken = os.Getenv("SLACK_BOT_API_TOKEN")
var godBotAPIToken = os.Getenv("SLACK_GOD_BOT_API_TOKEN")

func main() {
	slackClients := slackcr.NewSlackClients(botAPIToken, godBotAPIToken)
	slackCr := slackcr.New(slackClients, topics(slackClients))
	slackCr.Listen(port)
}
