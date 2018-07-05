package wi2guest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/topicutil"
)

var trigger = regexp.MustCompile(`(d|ドト|どと|doto|doutor) ?(w|わいふぁい|ワイファイ|wifi|WiFi|wi-fi|Wi-Fi)`)

var serverURL = os.Getenv("WI2_GUESTCODE_SERVER_URL")

func init() {
	if serverURL == "" {
		panic("WI2_GUESTCODE_SERVER_URL is not set")
	}
}

// Client keeps slack client to delete other's message.
type Client struct{}

// Talk is main Topic
func (Client) Talk(room chatroom.Room) chatroom.DidTalk {
	r := topicutil.WaitReceived(room)
	if !trigger.MatchString(r.Text) {
		return false
	}

	resp, err := http.Get(serverURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return true
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return true
	}

	room.Send(string(bs))
	return true
}
