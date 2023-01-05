package function

import (
	"fatsharkbot/src/util/cqhttp"
)

type SendSticker struct{}

var sendSticker *SendSticker

func (f *SendSticker) MatchesAwake(event *cqhttp.CqhttpEvent) bool {
	return true
}

func (f *SendSticker) MatchesListen(event *cqhttp.CqhttpEvent) bool {
	return true
}

func (f *SendSticker) MatchesPrivateAwake(event *cqhttp.CqhttpEvent) bool {
	return true
}

func (f *SendSticker) WorkAwake(event *cqhttp.CqhttpEvent) error {
	return nil
}

func (f *SendSticker) WorkPrivateAwake(event *cqhttp.CqhttpEvent) error {
	return nil
}

func (f *SendSticker) WorkListen(event *cqhttp.CqhttpEvent) error {
	return nil
}
