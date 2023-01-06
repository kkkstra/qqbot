package cqhttp

import (
	"bytes"
	"encoding/json"
	"fatsharkbot/src/config"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	getMsgUrl = fmt.Sprintf(config.C.Cqhttp.Addr + "/get_msg")
)

type ReqGetMessage struct {
	MessageId int32 `json:"message_id,omitempty"`
}

type ResGetMessage struct {
	Data DataGetMessage `json:"data,omitempty"`
}

type DataGetMessage struct {
	Group       bool      `json:"group,omitempty"`
	GroupId     int64     `json:"group_id,omitempty"`
	MessageId   int32     `json:"message_id,omitempty"`
	RealId      int32     `json:"real_id,omitempty"`
	MessageType string    `json:"message_type,omitempty"`
	Sender      sender    `json:"sender,omitempty"`
	Time        int32     `json:"time,omitempty"`
	Message     []message `json:"message,omitempty"`
	RawMessage  string    `json:"raw_message,omitempty"`
}

func GetMessage(msgId int32) (*ResGetMessage, error) {
	reqGetMessage := ReqGetMessage{
		MessageId: msgId,
	}
	getMessageJSON, _ := json.Marshal(reqGetMessage)
	resp, err := http.Post(getMsgUrl, "application/json", bytes.NewBuffer(getMessageJSON))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var res ResGetMessage
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}
