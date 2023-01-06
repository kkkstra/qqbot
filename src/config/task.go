package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Job struct {
	Tasks []Task `json:"tasks,omitempty"`
}

type Task struct {
	Spec    string   `json:"spec,omitempty"`
	Time    string   `json:"time,omitempty"`
	Msg     []string `json:"msg,omitempty"`
	Private []string `json:"private,omitempty"`
	Group   []string `json:"group,omitempty"`
}

var J *Job

func initTask() {
	taskFile := "task.json"

	file, err := os.Open(fmt.Sprintf("./env/task/%s", taskFile))
	if err != nil {
		log.Println(err)
		return
	}
	dec := json.NewDecoder(file)
	j := Job{}
	if err := dec.Decode(&j); err != nil {
		log.Println(err)
		return
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	J = &j
}
