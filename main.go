package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"encoding/json"
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/nlopes/slack"
	"net/http"
)

var slackIncomingWebhookURL = os.Getenv("SLACK_INCOMING_WEBHOOK_URL")
var port = os.Getenv("PORT")
var botAPIToken = os.Getenv("SLACK_BOT_API_TOKEN")

var api = slack.New(botAPIToken)

type incomingWebhook struct {
	Text string `json:"text"`
}

type received struct {
	text        string
	timestamp   string
	userName    string
	channelID   string
	channelName string
}

var crs = make(map[string]chatroom.Chatroom)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		riv := parseOutgoingWebhookQuery(body)
		if riv.userName == "slackbot" {
			return
		}
		cr, ok := crs[riv.channelID]
		if !ok {
			cr = chatroom.New(topics)
			crs[riv.channelID] = cr
		}
		// Pass the received message to Chatroom.
		cr.Flush(riv)
	})
	http.ListenAndServe(":"+port, nil)
}

func postToSlack(text string) {
	jsonStr, err := json.Marshal(incomingWebhook{Text: text})
	if err != nil {
		fmt.Println(err)
	}
	http.PostForm(slackIncomingWebhookURL, url.Values{"payload": {string(jsonStr)}})
}

func parseOutgoingWebhookQuery(body []byte) received {
	parsed, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println(err)
	}
	return received{
		text:        parsed["text"][0],
		timestamp:   parsed["timestamp"][0],
		userName:    parsed["user_name"][0],
		channelID:   parsed["channel_id"][0],
		channelName: parsed["channel_name"][0],
	}
}
