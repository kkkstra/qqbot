package cqhttp

import (
	"fatsharkbot/src/config"
	"strconv"
	"strings"
)

type CqhttpEvent struct {
	Time        int64     `json:"time,omitempty"`         // 时间戳
	SelfId      int64     `json:"self_id,omitempty"`      // 机器人QQ
	UserId      int64     `json:"user_id,omitempty"`      // 发送者QQ
	PostType    string    `json:"post_type,omitempty"`    // 上报类型
	NoticeType  string    `json:"notice_type,omitempty"`  // 通知类型
	MessageType string    `json:"message_type,omitempty"` // 消息类型
	SubType     string    `json:"sub_type,omitempty"`     // normal、anonymous、notice、poke
	MessageId   int32     `json:"message_id,omitempty"`   // 消息ID
	Message     []message `json:"message,omitempty"`      // 消息内容
	RawMessage  string    `json:"raw_message,omitempty"`  // 原始消息内容
	Font        int32     `json:"font,omitempty"`         // 字体
	GroupId     int64     `json:"group_id,omitempty"`     // 群号
	Sender      sender    `json:"sender,omitempty"`       // 发送人信息
	TargetId    int64     `json:"target_id,omitempty"`    // 被戳者 QQ 号
	Anonymous   anonymous `json:"anonymous,omitempty"`    // 匿名消息
	TempSource  int       `json:"temp_source,omitempty"`  // 临时消息来源
	OperatorId  int64     `json:"operator_id,omitempty"`
}

type message struct {
	Type string `json:"type,omitempty"`
	Data data   `json:"data,omitempty"`
}

type sender struct {
	UserId   int64  `json:"user_id,omitempty"`  // 发送者qq
	Nickname string `json:"nickname,omitempty"` // 昵称
	Sex      string `json:"sex,omitempty"`      // 性别，male、female、unknown
	Age      int32  `json:"age,omitempty"`      // 年龄
	GroupId  int64  `json:"group_id,omitempty"` // 临时群消息来源群号
	Card     string `json:"card,omitempty"`     // 群名片
	Area     string `json:"area,omitempty"`     // 地区
	Level    string `json:"level,omitempty"`    // 成员等级
	Role     string `json:"role,omitempty"`     // 角色，owner、admin、member
	Title    string `json:"title,omitempty"`    // 专属头衔
}

type data struct {
	QQ   string `json:"qq,omitempty"`   // CQ: at
	Text string `json:"text,omitempty"` // CQ: text
	File string `json:"file,omitempty"` // CQ: image
	Url  string `json:"url,omitempty"`  // CQ: image
	Id   string `json:"id,omitempty"`   // CQ: reply
}

type anonymous struct {
	Id   int64  `json:"id,omitempty"`   // 匿名用户id
	Name string `json:"name,omitempty"` // 匿名用户名称
	Flag string `json:"flag,omitempty"` // 匿名用户 flag, 在调用禁言 API 时需要传入
}

func (event *CqhttpEvent) AtBot() bool {
	return event.RawMessage == CodeAtBotSpace ||
		event.RawMessage == CodeAtBot ||
		event.RawMessage == RawAtBot ||
		event.RawMessage == RawAtBotSpace
}

func (event *CqhttpEvent) InteractWithBot() bool {
	for _, v := range event.Message {
		if v.Type == "at" && v.Data.QQ == config.C.Bot.Qq {
			return true
		}
	}
	//if strings.Contains(event.RawMessage, RawAtBot) || strings.Contains(event.RawMessage, CodeAtBot) {
	//	return true
	//}
	return false
}

func (event *CqhttpEvent) PokeBot() bool {
	return event.SubType == "poke" && strconv.FormatInt(event.TargetId, 10) == config.C.Bot.Qq
}

func (event *CqhttpEvent) GroupIncrease() bool {
	return event.NoticeType == "group_increase"
}

func (event *CqhttpEvent) TextContainsAny(keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(event.RawMessage, keyword) {
			return true
		}
	}
	return false
}

func (event *CqhttpEvent) TextHasPrefix(keyword string) bool {
	for _, v := range event.Message {
		if v.Type == "text" && strings.HasPrefix(strings.TrimPrefix(v.Data.Text, " "), keyword) {
			return true
		}
	}
	//rawMessage := strings.TrimPrefix(event.RawMessage, " ")
	//if strings.HasPrefix(strings.TrimPrefix(rawMessage, RawAtBot), keyword) ||
	//	strings.HasPrefix(strings.TrimPrefix(rawMessage, RawAtBotSpace), keyword) {
	//	return true
	//}
	return false
}
