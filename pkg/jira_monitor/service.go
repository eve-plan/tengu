package jira_monitor

import (
	"sync"

	"github.com/robfig/cron/v3"

	"tengu/pkg/util"
)

var wait sync.WaitGroup

// TODO log
func StartService() error {
	defer func() {
		if err := recover(); err != nil {
			util.PanicHandler()
		}
	}()
	if err := LoadUserInfo(); err != nil {
		return err
	}
	InitCorn()
	wait.Add(1)
	wait.Wait()
	return nil
}

func InitCorn() {
	go func() {
		c := cron.New()
		defer c.Run()
		cornHandler(c)
	}()
}

func cornHandler(c *cron.Cron) {
	_, err := c.AddFunc("@every 1m", MonitorBugs)
	if err != nil {
		panic(err)
	}
}
