package jira_monitor

import (
	"log"
	"os"
	"strings"

	"tengu/pkg/util"
)

const (
	bugPath = "/rest/api/2/search"
)

var host string
var authorization string

func LoadUserInfo() error {
	authorization = ""
	bytes, err := os.ReadFile("./jira_user_info.conf")
	if err != nil {
		return err
	}
	info := string(bytes)
	split := strings.Split(info, "|")
	host = split[0]
	authorization = split[1]
	return nil
}

func GetBugs() {
	urlStr := host + bugPath
	body, err := util.HttpGet(urlStr, map[string]string{
		"jql": "issuetype = 故障 AND resolution = Unresolved AND assignee in (currentUser()) order by updated DESC",
	}, map[string]string{"Authorization": authorization})
	if err != nil {
		log.Print(err)
		return
	}
	log.Println(string(body))
}
