package function

import (
	"fatsharkbot/src/util/cqhttp"
)

type Function interface {
	// 被@时的匹配规则
	MatchesAwake(*cqhttp.CqhttpEvent) bool
	// 未@时的匹配规则
	MatchesListen(*cqhttp.CqhttpEvent) bool
	// 私聊匹配
	MatchesPrivateAwake(*cqhttp.CqhttpEvent) bool
	// 触发动作
	WorkPrivateAwake(*cqhttp.CqhttpEvent) error
	WorkAwake(*cqhttp.CqhttpEvent) error
	WorkListen(*cqhttp.CqhttpEvent) error
}

// sorted by priority
var funcs = []Function{
	sendSticker,
	autoReply,
}

func Awake(event *cqhttp.CqhttpEvent) error {
	for _, f := range funcs {
		if f.MatchesAwake(event) {
			return f.WorkAwake(event)
		}
	}
	return nil
}

func Listen(event *cqhttp.CqhttpEvent) error {
	for _, f := range funcs {
		if f.MatchesListen(event) {
			return f.WorkListen(event)
		}
	}
	return nil
}

func PrivateAwake(event *cqhttp.CqhttpEvent) error {
	for _, f := range funcs {
		if f.MatchesPrivateAwake(event) {
			return f.WorkPrivateAwake(event)
		}
	}
	return nil
}
