package reply

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/util/cqhttp"
	"log"
)

type AutoReply struct{}

var (
	AutoReplyFunc    *AutoReply
	groupIncreaseMsg = "欢迎新人，看看精华消息，看看群公告，有问题如果方便的话可以发在群里，众多优秀学长学姐可以一起解答，记得改备注。\n\n如果想和肥肥鲨聊天，又不知如何操作，不如@肥肥鲨发送 \"help\" 或 \"帮助\" 试试！"
)

func SendDefaultPrivateMsg(event *cqhttp.CqhttpEvent) error {
	msg := "你好qwq，我是肥肥鲨bot！" + cqhttp.GetRandomCodeFace()
	err := event.SendPrivateMsg(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SendDefaultGroupMsg(event *cqhttp.CqhttpEvent) error {
	err := event.SendRandGroupMsgWithFace(config.R.DefaultMsg, cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (f *AutoReply) MatchesAwake(event *cqhttp.CqhttpEvent) bool {
	return true
}

func (f *AutoReply) MatchesListen(event *cqhttp.CqhttpEvent) bool {
	return false
}

func (f *AutoReply) MatchesPrivateAwake(event *cqhttp.CqhttpEvent) bool {
	return true
}

func (f *AutoReply) WorkAwake(event *cqhttp.CqhttpEvent) error {
	for _, reply := range config.R.Msg {
		if event.TextContainsAny(reply.Keywords) {
			err := event.SendRandGroupMsgWithFace(reply.Messages, cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	err := event.SendRandGroupMsgWithFace(config.R.DefaultMsg, cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (f *AutoReply) WorkPrivateAwake(event *cqhttp.CqhttpEvent) error {
	for _, reply := range config.R.Msg {
		if event.TextContainsAny(reply.Keywords) {
			err := event.SendRandPrivateMsgWithFace(reply.Messages, cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	err := event.SendRandPrivateMsgWithFace(config.R.DefaultMsg, cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (f *AutoReply) WorkListen(event *cqhttp.CqhttpEvent) error {
	return nil
}

func SendMsgWhenGroupIncrease(event *cqhttp.CqhttpEvent) error {
	err := event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + groupIncreaseMsg + cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
