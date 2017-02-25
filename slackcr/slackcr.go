package slackcr

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/nlopes/slack"
	"net/http"
)

// SlackClients have all of client for Slack.
type SlackClients struct {
	Friends *slack.Client
	God *slack.Client
}

// Client keeps the clients for Slack Bot and methods to operate them.
type Client struct {
	slackClients SlackClients
	topics []chatroom.Topic
}

// NewSlackClients cretes SlackClients
func NewSlackClients(botAPIToken string, godBotAPIToken string) SlackClients {
	return SlackClients{
		Friends: slack.New(botAPIToken),
		God: slack.New(godBotAPIToken),
	}
}

// New creates Client.
func New(slackClients SlackClients, topics []chatroom.Topic) Client {
	return Client{
		slackClients: slackClients,
		topics: topics,
	}
}

// Received type have request of Slack Outgoing Webhook.
type Received struct {
	Text        string
	Timestamp   string
	UserName    string
	ChannelID   string
	ChannelName string
}

// BotProfile express the name and icon(Emoji) of one Bot.
type BotProfile struct {
	UserName string
	IconEmoji string
}

var botProfiles = []BotProfile{
	BotProfile{
		UserName: "God",
		IconEmoji: ":god:",
	},
}

var crs = make(map[string]chatroom.Chatroom)

// Listen listens the outgoing webhook integration of Slack.
func (client *Client) Listen(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		riv := parseOutgoingWebhookQuery(body)
		if riv.UserName == "slackbot" {
			return
		}
		cr, ok := crs[riv.ChannelID]
		if !ok {
			cr = chatroom.New(client.topics)
			crs[riv.ChannelID] = cr

			go client.sender(riv.ChannelID)
		}
		// Pass the received message to Chatroom.
		cr.Flush(riv)
	})
	http.ListenAndServe(":"+port, nil)
}

func (client *Client) sender(channelID string) {
	cr, ok := crs[channelID]
	if !ok {
		fmt.Println("ERR: the chatroom is not found. It's must be.")
		return
	}
	for {
		text := cr.WaitSentTextMsg()
		client.postToSlack(channelID, text)
	}
}

func (client *Client) postToSlack(channelID string, text string) {
	_, _, err := client.slackClients.God.PostMessage(channelID, text, slack.PostMessageParameters{Username: "God", IconEmoji: ":kemono_friends:"})
	if err != nil {
		fmt.Println(err)
	}
}

func parseOutgoingWebhookQuery(body []byte) Received {
	parsed, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println(err)
	}
	return Received{
		Text:        parsed["text"][0],
		Timestamp:   parsed["timestamp"][0],
		UserName:    parsed["user_name"][0],
		ChannelID:   parsed["channel_id"][0],
		ChannelName: parsed["channel_name"][0],
	}
}

// SlackClients returns all of Slack Client the client have.
func (client *Client) SlackClients() SlackClients {
	return client.slackClients
}
