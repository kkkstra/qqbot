package cron

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/util/cqhttp"
	"github.com/robfig/cron/v3"
	"log"
	"strconv"
	"time"
)

func TimedMsg() {
	go func() {
		for _, job := range config.J.Tasks {
			t, err := time.Parse("15:04:05", job.Time)
			if err != nil {
				log.Println(nil)
			}
			now := time.Now()
			realT := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), now.Location())
			if now.Sub(realT) > 0 && now.Sub(realT) < 3*time.Second {
				for _, groupId := range job.Group {
					id, _ := strconv.ParseInt(groupId, 10, 64)
					err := (&cqhttp.CqhttpEvent{GroupId: id}).SendRandGroupMsgWithFace(job.Msg, cqhttp.GetRandomCodeFace())
					if err != nil {
						log.Println(err)
					}
				}
				for _, userId := range job.Private {
					id, _ := strconv.ParseInt(userId, 10, 64)
					err := (&cqhttp.CqhttpEvent{UserId: id}).SendRandPrivateMsgWithFace(job.Msg, cqhttp.GetRandomCodeFace())
					if err != nil {
						log.Println(err)
					}
				}
				time.Sleep(3 * time.Second)
			}
		}
	}()
}

func InitCron() {
	c := cron.New(cron.WithSeconds())

	for _, job := range config.J.Tasks {

		_, err := c.AddFunc(job.Spec, TimedMsg)
		if err != nil {
			log.Println(err)
			return
		}
	}

	c.Start()
}
