package function

import (
	"fatsharkbot/src/util/cqhttp"
	"log"
)

func PokePoke(event *cqhttp.CqhttpEvent) error {
	if event.GroupId != 0 {
		err := event.SendGroupMsg(cqhttp.GetCodePoke(event.UserId))
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		err := event.SendPrivateMsg(cqhttp.GetCodePoke(event.UserId))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
