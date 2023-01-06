package controller

import (
	"fatsharkbot/src/function"
	"fatsharkbot/src/function/poke"
	"fatsharkbot/src/function/reply"
	"fatsharkbot/src/util/cqhttp"
	"github.com/gin-gonic/gin"
)

func EventHandler(c *gin.Context) {
	event := new(cqhttp.CqhttpEvent)
	c.Bind(event)
	switch event.PostType {
	case "message":
		messageHandler(event)
	case "notice":
		noticeHandler(event)
	}
}

func messageHandler(event *cqhttp.CqhttpEvent) {
	//fmt.Println(event)
	if event.MessageType == "group" {
		switch {
		// simply @bot
		case event.AtBot():
			_ = reply.SendDefaultGroupMsg(event)
		// @bot and say something
		case event.InteractWithBot():
			_ = function.Awake(event)
		// listen
		default:
			_ = function.Listen(event)
		}
	} else {
		_ = function.PrivateAwake(event)
	}
}

func noticeHandler(event *cqhttp.CqhttpEvent) {
	// 拍一拍
	if event.PokeBot() {
		_ = poke.PokePoke(event)
	}
}
