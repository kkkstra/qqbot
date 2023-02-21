package sticker

import (
	"fatsharkbot/src/util/cqhttp"
	"log"
	"strings"
)

type CollectSticker struct{}

var (
	CollectStickerFunc *CollectSticker
	collectStickerList = map[int64]string{}
)

func (f *CollectSticker) MatchesAwake(event *cqhttp.CqhttpEvent) bool {
	return event.TextHasPrefix("收藏")
}

func (f *CollectSticker) MatchesListen(event *cqhttp.CqhttpEvent) bool {
	return event.TextHasPrefix("收藏")
}

func (f *CollectSticker) MatchesPrivateAwake(event *cqhttp.CqhttpEvent) bool {
	return event.TextHasPrefix("收藏")
}

func (f *CollectSticker) WorkAwake(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *CollectSticker) WorkListen(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *CollectSticker) WorkPrivateAwake(event *cqhttp.CqhttpEvent) error {
	for _, v := range event.Message {
		if v.Type == "text" && strings.Contains(v.Data.Text, "tag=") {
			i := strings.Index(v.Data.Text, "tag=")
			tag := v.Data.Text[i+4:]
			tag = strings.TrimPrefix(tag, " ")
			tag = strings.TrimSuffix(tag, " ")
			if tag == "" {
				err := event.SendPrivateMsg("收藏图片的标签不能为空哦！" + cqhttp.GetRandomCodeFace())
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			}
			err := event.SendPrivateMsg("收到！快把想收藏的 \"" + tag + "\" 图片告诉肥肥鲨吧！" + cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			collectStickerList[event.UserId] = tag
			return nil
		}
	}
	// 未找到tag
	err := event.SendPrivateMsg("收藏图片的正确格式是\"收藏 tag=标签\"哦，尝试修改一下再告诉肥肥鲨吧qwq！" + cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (f *CollectSticker) WorkGroup(event *cqhttp.CqhttpEvent) error {
	for _, v := range event.Message {
		if v.Type == "text" && strings.Contains(v.Data.Text, "tag=") {
			i := strings.Index(v.Data.Text, "tag=")
			tag := v.Data.Text[i+4:]
			tag = strings.TrimPrefix(tag, " ")
			tag = strings.TrimSuffix(tag, " ")
			if tag == "" {
				err := event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "收藏图片的标签不能为空哦！" + cqhttp.GetRandomCodeFace())
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			}
			err := event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "收到！快把想收藏的 \"" + tag + "\" 图片告诉肥肥鲨吧！" + cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			collectStickerList[event.UserId] = tag
			return nil
		}
	}
	// 未找到tag
	err := event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "收藏图片的正确格式是\"收藏 tag=标签\"哦，尝试修改一下再告诉肥肥鲨吧qwq！" + cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
