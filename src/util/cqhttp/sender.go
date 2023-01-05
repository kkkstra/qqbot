package cqhttp

import (
	"bytes"
	"encoding/json"
	"fatsharkbot/src/config"
	"fatsharkbot/src/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var (
	groupMsgUrl   = fmt.Sprintf(config.C.Cqhttp.Addr + "/send_group_msg")
	privateMsgUrl = fmt.Sprintf(config.C.Cqhttp.Addr + "/send_private_msg")
)

type ReqSendGroupMsg struct {
	GroupId    string `json:"group_id,omitempty"`
	Message    string `json:"message,omitempty"`
	AutoEscape bool   `json:"auto_escape,omitempty"`
}

type ReqSendPrivateMsg struct {
	UserId     string `json:"user_id,omitempty"`
	GroupId    string `json:"group_id,omitempty"`
	Message    string `json:"message,omitempty"`
	AutoEscape bool   `json:"auto_escape,omitempty"`
}

type ResSendMsg struct {
	Status  string `json:"status,omitempty"`
	Retcode int    `json:"retcode,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Wording string `json:"wording,omitempty"`
	Data    struct {
		MessageId int `json:"message_id,omitempty"`
	} `json:"data,omitempty"`
	Echo string `json:"echo,omitempty"`
}

func (event *CqhttpEvent) SendGroupMsg(msg string) error {
	groupId := strconv.FormatInt(event.GroupId, 10)
	groupMsg := ReqSendGroupMsg{
		GroupId:    groupId,
		Message:    msg,
		AutoEscape: false,
	}
	groupMsgJSON, _ := json.Marshal(groupMsg)
	resp, err := http.Post(groupMsgUrl, "application/json", bytes.NewBuffer(groupMsgJSON))
	if err != nil {
		log.Println(err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	var res ResSendMsg
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (event *CqhttpEvent) SendPrivateMsg(msg string) error {
	userId := strconv.FormatInt(event.UserId, 10)
	privateMsg := ReqSendPrivateMsg{
		UserId:     userId,
		Message:    msg,
		AutoEscape: false,
	}
	privateMsgJSON, _ := json.Marshal(privateMsg)
	resp, err := http.Post(privateMsgUrl, "application/json", bytes.NewBuffer(privateMsgJSON))
	if err != nil {
		log.Println(err)
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	var res ResSendMsg
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (event *CqhttpEvent) SendRandGroupMsg(msgList []string) error {
	msg := util.RandString(msgList)
	err := event.SendGroupMsg(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (event *CqhttpEvent) SendRandGroupMsgWithFace(msgList []string, face string) error {
	msg := util.RandString(msgList) + face
	err := event.SendGroupMsg(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (event *CqhttpEvent) SendRandPrivateMsg(msgList []string) error {
	msg := util.RandString(msgList)
	err := event.SendPrivateMsg(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (event *CqhttpEvent) SendRandPrivateMsgWithFace(msgList []string, face string) error {
	msg := util.RandString(msgList) + face
	err := event.SendPrivateMsg(msg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
