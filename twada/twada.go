package twada

import (
	"regexp"

	"github.com/acomagu/chatroom-go/chatroom"
	"github.com/acomagu/gcf-slack-bot/topicutil"
)

var trigger = regexp.MustCompile(`(テスト|てすと)(書いてない|かいてない|かかなきゃ|書かなきゃ|だる|しんど|無理|むり|不可|やりたく|したく)`)

// Client keeps slack client to delete other's message.
type Client struct{}

// Talk is main Topic
func (Client) Talk(room chatroom.Room) chatroom.DidTalk {
	r := topicutil.WaitReceived(room)
	if !trigger.MatchString(r.Text) {
		return false
	}

	room.Send("```\n" + `　　　　 ,、,,,、,,, 
　　　 _,,;' '" '' ;;,, 
　　（rヽ,;''""''゛゛;,ﾉｒ）　　　　 
　　 ,; i ___　、___iヽ゛;,　　テスト書いてないとかお前それ@t_wadaの前でも同じ事言えんの？
　 ,;'''|ヽ・〉〈・ノ |ﾞ ';, 
　 ,;''"|　 　▼　　 |ﾞ゛';, 
　 ,;''　ヽ ＿人＿  /　,;'_ 
 ／ｼ、    ヽ  ⌒⌒  /　 ﾘ　＼ 
|　　 "ｒ,,｀"'''ﾞ´　　,,ﾐ| 
|　　 　 ﾘ、　　　　,ﾘ　　 | 
|　　i 　゛ｒ、ﾉ,,ｒ" i _ | 
|　　｀ー――-----------┴ ⌒´ ） 
（ヽ  _____________ ,, ＿´） 
 （_⌒_______________ ,, ィ 
     T                 |
     |                 |` + "\n```")
	return true
}
