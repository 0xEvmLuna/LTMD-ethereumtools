package bot

import (
	"fmt"

	cron "github.com/robfig/cron/v3"
)

// Define common fixed time
const (
	// Time interval for second execution
	Second = "@every 1s"

	// Time interval for hourly execution
	Hour = "@every 1h"

	// Time interval for day execution
	Day = "@daily"

	// Time interval for week execution
	Week = "@weekly"

	// Time interval for month execution
	Month = "@monthly"
)

type Timer interface {
	NewTimerTask() *TimerTask
	AddTask()
}

type TimerTask struct {
	Cron    *cron.Cron
	Time    string
	Release chan struct{}
}

func NewTimerTask() *TimerTask {
	return &TimerTask{
		Cron:    cron.New(),
		Release: make(chan struct{}),
	}
}

func (t *TimerTask) AddTask(spec string) {
	t.Cron.AddFunc(spec, func() { fmt.Printf("Run: 任务开始%d\n") })
}

func (t *TimerTask) AddJob(spec string, callback func()) {
	t.Cron.AddFunc(spec, callback)
}

func (t *TimerTask) Start() {
	t.Cron.Start()
	<-t.Release
}

func (t *TimerTask) Stop() {
	t.Cron.Stop()
	close(t.Release)
}
