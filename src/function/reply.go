package function

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/util/cqhttp"
	"log"
)

type AutoReply struct{}

var autoReply *AutoReply

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
	return true
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
	//for _, reply := range config.R.RepliesWithoutAt {
	//	if event.TextContainsAny(reply.Keywords) {
	//		err := event.SendGroupMsgList(reply.Message)
	//		if err != nil {
	//			log.Logger.Error(err)
	//		}
	//		return nil
	//	}
	//}
	return nil
}
