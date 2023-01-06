package sticker

import (
	"bufio"
	"fatsharkbot/src/config"
	"fatsharkbot/src/model"
	"fatsharkbot/src/util/cqhttp"
	"io"
	"log"
	"net/http"
	"os"
)

type SaveSticker struct{}

var SaveStickerFunc *SaveSticker

func (f *SaveSticker) MatchesAwake(event *cqhttp.CqhttpEvent) bool {
	return f.Matches(event)
}

func (f *SaveSticker) MatchesListen(event *cqhttp.CqhttpEvent) bool {
	return f.Matches(event)
}

func (f *SaveSticker) MatchesPrivateAwake(event *cqhttp.CqhttpEvent) bool {
	return f.Matches(event)
}

func (f *SaveSticker) Matches(event *cqhttp.CqhttpEvent) bool {
	_, ok := collectStickerList[event.UserId]
	if !ok {
		return false
	}
	for _, v := range event.Message {
		if v.Type == "image" {
			return true
		}
	}
	return false
}

func (f *SaveSticker) WorkAwake(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *SaveSticker) WorkListen(event *cqhttp.CqhttpEvent) error {
	return f.WorkGroup(event)
}

func (f *SaveSticker) WorkPrivateAwake(event *cqhttp.CqhttpEvent) error {
	tag, _ := collectStickerList[event.UserId]
	delete(collectStickerList, event.UserId)
	for _, v := range event.Message {
		if v.Type == "image" {
			file, url := v.Data.File, v.Data.Url
			savePath := config.C.App.ImagePath + file
			err := saveImage(url, savePath)
			if err != nil {
				log.Println(err)
				return err
			}
			m := model.GetModel(&model.StickerModel{})
			s := &model.Sticker{
				Tag:  tag,
				File: file,
				Url:  "file://" + savePath,
			}
			err = m.(*model.StickerModel).CreateSticker(s)
			if err != nil {
				log.Println(err)
				return err
			}
			err = event.SendPrivateMsg("肥肥鲨收到标签为 \"" + tag + "\" 的图片啦！" + cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	return nil
}

func (f *SaveSticker) WorkGroup(event *cqhttp.CqhttpEvent) error {
	tag, _ := collectStickerList[event.UserId]
	delete(collectStickerList, event.UserId)
	for _, v := range event.Message {
		if v.Type == "image" {
			file, url := v.Data.File, v.Data.Url
			savePath := config.C.App.ImagePath + file
			err := saveImage(url, savePath)
			if err != nil {
				log.Println(err)
				return err
			}
			m := model.GetModel(&model.StickerModel{})
			s := &model.Sticker{
				Tag:  tag,
				File: file,
				Url:  "file://" + savePath,
			}
			err = m.(*model.StickerModel).CreateSticker(s)
			if err != nil {
				log.Println(err)
				return err
			}
			err = event.SendGroupMsg(cqhttp.GetCodeAt(event.UserId) + "肥肥鲨收到标签为 \"" + tag + "\" 的图片啦！" + cqhttp.GetRandomCodeFace())
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		}
	}
	return nil
}

func saveImage(url, path string) error {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = io.Copy(writer, reader)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
