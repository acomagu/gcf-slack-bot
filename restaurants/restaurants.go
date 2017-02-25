package restaurants

import (
	"math/rand"
	"regexp"
	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/nlopes/slack"
	"github.com/acomagu/gcf-slack-bot/topicutil"
)

type isAccepted bool

// Client is restaurants slack client struct
type Client struct {
	bot *slack.Client
}

var mealTrigger = regexp.MustCompile("(めし|飯|ごはん|ご飯)ルーレット")
var majorRestaurasts = []string{
	"二郎",
	"幸楽苑",
	"お好み焼き廣",
	"あおやま",
	"デニーズ",
	"ガスト",
	"中華しせんや",
	"風街亭",
	"One's home",
	"ハジャイ",
	"ビストロビーグルシェフ",
	"居酒屋もみじ",
	"tipu",
	"会津バーガーラッキースマイル",
	"沙羅屋",
	"皆川食肉店",
	"三番山下",
	"ほるもん道場",
	"伝八",
	"めでたいや",
	"ココス",
	"すき家",
	"吉野家",
	"よどや",
	"麺屋ごんちゃん",
	"丼丸",
	"とん八",
	"重慶飯店",
	"マクドナルド",
}
var minorRestaurants = []string{"学食", "自宅", "絶食", "物乞い"}

var positiveReacts = []*regexp.Regexp{
	regexp.MustCompile(`あり`),
	regexp.MustCompile(`いいね`),
	regexp.MustCompile(`(よ|良)さ`),
	regexp.MustCompile(`(よ|良)い`),
	regexp.MustCompile(`わかる`),
	regexp.MustCompile(`おお`),
}
var negativeReacts = []*regexp.Regexp{
	regexp.MustCompile(`(な|無)し`),
	regexp.MustCompile(`もう一回`),
	regexp.MustCompile(`(な|無)い`),
	regexp.MustCompile(`無理`),
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
	if !mealTrigger.MatchString(r.Text) {
		return false
	}

	for {
		room.Send("今日のごはんは...")
		if isTodayMajor() {
			room.Send("じゃーん! *" + todaysMajorRestaurant() + "*!")
			if majorReact(room) {
				return true
			}
		} else {
			room.Send("じゃーん! *" + todaysMinorRestaurant() + "*!")
			if minorReact(room) {
				return true
			}
		}
	}
}

func majorReact(room chatroom.Room) isAccepted {
	r := topicutil.WaitReceived(room)
	if isPositiveReaction(r.Text) {
		room.Send(choose([]string{
			"ありがとう!",
			"いいでしょ!",
			"すごーい!",
		}))
		return true
	} else if isNegativeReaction(r.Text) {
		room.Send(choose([]string{
			"はぁ!? きれそう",
			"もう疲れたよ...",
			"お気に召さなかったですか...? もう一度やらせてください!",
			"すみません、無能で...",
			"もう一回やってもいいですか?",
			"すみません...もう一回させてください...!",
		}))
		return false
	}
	return true
}

func minorReact(room chatroom.Room) isAccepted {
	r := topicutil.WaitReceived(room)
	if isPositiveReaction(r.Text) {
		room.Send(choose([]string{"え...ひくわ...", "まじか...", "(まじか...)", "ふーん?", "すごい...ね?", "へぇ...そういうのが好きなの?"}))
		return true
	} else if isNegativeReaction(r.Text) {
		room.Send(choose([]string{"贅沢だな! もう...", "ふふ...もう、しかたないなぁ", "贅沢言わないでくれる? はぁ..."}))
		return false
	}
	return true
}

func isNegativeReaction(str string) bool {
	return topicutil.MatchAny(negativeReacts, str)
}

func isPositiveReaction(str string) bool {
	return topicutil.MatchAny(positiveReacts, str)
}

func isTodayMajor() bool {
	return rand.Intn(5) != 0
}

func todaysMajorRestaurant() string {
	return choose(majorRestaurasts)
}

func todaysMinorRestaurant() string {
	return choose(minorRestaurants)
}

func choose(arr []string) string {
	return arr[rand.Intn(len(arr))]
}
