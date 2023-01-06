package sticker

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/model"
	"fatsharkbot/src/util"
	"fatsharkbot/src/util/cqhttp"
	"fmt"
	"log"
	"strings"
)

type SendSticker struct{}

var SendStickerFunc *SendSticker

func (f *SendSticker) MatchesAwake(event *cqhttp.CqhttpEvent) bool {
	return event.TextHasPrefix("发")
}

func (f *SendSticker) MatchesListen(event *cqhttp.CqhttpEvent) bool {
	return event.TextHasPrefix("发")
}

func (f *SendSticker) MatchesPrivateAwake(event *cqhttp.CqhttpEvent) bool {
	return event.TextHasPrefix("发")
}

func (f *SendSticker) WorkAwake(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *SendSticker) WorkListen(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *SendSticker) WorkPrivateAwake(event *cqhttp.CqhttpEvent) error {
	for _, v := range event.Message {
		if v.Type == "text" && strings.Contains(v.Data.Text, "发") {
			i := strings.Index(v.Data.Text, "发")
			tag := v.Data.Text[i+3:]
			m := model.GetModel(&model.StickerModel{})
			res, err := m.(*model.StickerModel).GetStickerByTag(tag)
			if err != nil {
				log.Println(err)
				return err
			}
			resNum := len(*res)
			if resNum == 0 {
				// 未找到结果
				err := event.SendPrivateMsg("很抱歉，肥肥鲨没有找到 \"" + tag + "\" 相关图片哦TvT" + cqhttp.GetRandomCodeFace())
				if err != nil {
					log.Println(err)
					return err
				}
			} else {
				n := util.RandInt32(int32(resNum)) // 要发送的编号
				msg := fmt.Sprintf("[CQ:image,file=%s,url=%s]", (*res)[n].File, (*res)[n].Url)
				err := event.SendPrivateMsg(msg)
				if err != nil {
					log.Println(err)
					return err
				}
			}
			return nil
		}
	}
	return nil
}

func (f *SendSticker) WorkGroup(event *cqhttp.CqhttpEvent) error {
	for _, v := range event.Message {
		if v.Type == "text" && strings.Contains(v.Data.Text, "发") {
			i := strings.Index(v.Data.Text, "发")
			tag := v.Data.Text[i+3:]
			m := model.GetModel(&model.StickerModel{})
			res, err := m.(*model.StickerModel).GetStickerByTag(tag)
			if err != nil {
				log.Println(err)
				return err
			}
			resNum := len(*res)
			if resNum == 0 {
				// 未找到结果
				err := event.SendGroupMsg("很抱歉，肥肥鲨没有找到 \"" + tag + "\" 相关图片哦TvT" + cqhttp.GetRandomCodeFace())
				if err != nil {
					log.Println(err)
					return err
				}
			} else {
				n := util.RandInt32(int32(resNum)) // 要发送的编号
				msg := fmt.Sprintf("[CQ:image,file=%s,subType=0,url=%s]", (*res)[n].File, config.C.App.ImagePath+(*res)[n].File)
				err := event.SendGroupMsg(msg)
				if err != nil {
					log.Println(err)
					return err
				}
			}
			return nil
		}
	}
	return nil
}
