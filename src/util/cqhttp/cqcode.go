package cqhttp

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/util"
	"fmt"
	"strconv"
)

var (
	CodeFace    = "[CQ:face,id=%s]"
	CodeRecord  = "[CQ:record,file=%s]"
	CodeAt      = "[CQ:at,qq=%s]"
	CodeAtSpace = "[CQ:at,qq=%s] "
	CodePoke    = "[CQ:poke,qq=%s]"

	CodeAtBot      = fmt.Sprintf(CodeAt, config.C.Bot.Qq)
	CodeAtBotSpace = fmt.Sprintf(CodeAtSpace, config.C.Bot.Qq)
	RawAtBot       = fmt.Sprintf("@%s", config.C.Bot.Name)
	RawAtBotSpace  = fmt.Sprintf("@%s ", config.C.Bot.Name)
)

func GetRandomCodeFace() string {
	faceId := util.RandInt32(222)
	return fmt.Sprintf(CodeFace, strconv.Itoa(int(faceId)))
}

func GetCodePoke(qq int64) string {
	return fmt.Sprintf(CodePoke, strconv.Itoa(int(qq)))
}

func GetCodeAt(qq int64) string {
	return fmt.Sprintf(CodeAt, strconv.Itoa(int(qq)))
}
