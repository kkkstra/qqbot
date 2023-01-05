package cqhttp

import (
	"fatsharkbot/src/config"
	"fmt"
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
