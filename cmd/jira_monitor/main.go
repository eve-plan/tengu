package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"

	"tengu/pkg/jira_monitor"
)

const (
	SERVICE = "service"
	TASK    = "task"
)

func init() {
	fp, _ := os.OpenFile("./jira_monitor.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	log.SetOutput(fp)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		subProcess([]string{os.Args[0], SERVICE})
		os.Exit(0)
	}
	switch os.Args[1] {
	case SERVICE:
		for {
			cmd := subProcess([]string{os.Args[0], TASK})
			log.Println("task start")
			err := cmd.Wait()
			if err != nil {
				log.Fatal(err.Error())
			}
			time.Sleep(time.Second * 10)
			log.Println("service continue")
			continue
		}
	case TASK:
		log.Println("task running")
		if err := jira_monitor.StartService(); err != nil {
			log.Fatal(err)
		}
		log.Println("task exit")

	default:
		must(errors.New("unknown args"))
	}
}

func subProcess(args []string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		must(err)
	}
	return cmd
}
