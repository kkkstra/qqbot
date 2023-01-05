package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type reply struct {
	DefaultMsg []string `json:"default_msg"`
	Msg        []msg    `json:"msg"`
}

type msg struct {
	Keywords []string `json:"keywords"`
	Messages []string `json:"messages"`
}

var R *reply

func initReply() {
	replyFile := "reply.json"

	//Rep := reply{}
	//key1 := []string{"a", "b", "c"}
	//key2 := []string{"d", "e", "f"}
	//msg1 := []string{"1", "2", "3"}
	//msg2 := []string{"4", "5", "6"}
	//
	//Rep.Msg = append(Rep.Msg, msg{
	//	Keywords: key1,
	//	Messages: msg1,
	//})
	//Rep.Msg = append(Rep.Msg, msg{
	//	Keywords: key2,
	//	Messages: msg2,
	//})
	//
	//file, err := os.Create(fmt.Sprintf("./env/msg/%s", replyFile))
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//enc := json.NewEncoder(file)
	////for _, v := range Rep.Msg {
	//err = enc.Encode(&Rep)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	////}
	//err = file.Close()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	file, err := os.Open(fmt.Sprintf("./env/msg/%s", replyFile))
	if err != nil {
		log.Println(err)
	}
	dec := json.NewDecoder(file)
	r := reply{}
	if err := dec.Decode(&r); err != nil {
		log.Println(err)
		return
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
	}
	R = &r
}
