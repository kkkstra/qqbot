package sticker

import (
	"fatsharkbot/src/model"
	"fatsharkbot/src/util/cqhttp"
	"log"
	"strconv"
	"strings"
)

type DeleteSticker struct{}

var DeleteStickerFunc *DeleteSticker

func (f *DeleteSticker) MatchesAwake(event *cqhttp.CqhttpEvent) bool {
	return f.Matches(event)
}

func (f *DeleteSticker) MatchesListen(event *cqhttp.CqhttpEvent) bool {
	return f.Matches(event)
}

func (f *DeleteSticker) MatchesPrivateAwake(event *cqhttp.CqhttpEvent) bool {
	return f.Matches(event)
}

func (f *DeleteSticker) Matches(event *cqhttp.CqhttpEvent) bool {
	hasReply, hasDelete := false, false
	for _, v := range event.Message {
		if v.Type == "reply" {
			hasReply = true
		}
		if v.Type == "text" && strings.Contains(v.Data.Text, "删除") {
			hasDelete = true
		}
		if hasReply && hasDelete {
			return true
		}
	}
	return false
}

func (f *DeleteSticker) WorkAwake(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *DeleteSticker) WorkListen(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *DeleteSticker) WorkPrivateAwake(event *cqhttp.CqhttpEvent) error {
	return nil
}

func (f *DeleteSticker) WorkGroup(event *cqhttp.CqhttpEvent) error {
	if event.Sender.Role == "member" {
		err := event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "只有管理员才能删除图片哦" + cqhttp.GetRandomCodeFace())
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
	var replyId int
	for _, v := range event.Message {
		if v.Type == "reply" {
			replyId, _ = strconv.Atoi(v.Data.Id)
			break
		}
	}
	originMsg, err := cqhttp.GetMessage(int32(replyId))
	if err != nil {
		log.Println(err)
		return err
	}
	msgData := originMsg.Data
	for _, v := range msgData.Message {
		if v.Type == "image" {
			file := v.Data.File
			m := model.GetModel(&model.StickerModel{})
			res, _ := m.(*model.StickerModel).GetStickerByFile(file)
			if len(*res) == 0 {
				err = event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "未找到图片，找到要删除的图片再来找肥肥鲨吧！" + cqhttp.GetRandomCodeFace())
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			}
			err := m.(*model.StickerModel).DeleteSticker(file)
			if err != nil {
				log.Println(err)
				return err
			}
			err = event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "肥肥鲨已成功删除照片！" + cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	err = event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "未找到图片，找到要删除的图片再来找肥肥鲨吧！" + cqhttp.GetRandomCodeFace())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
