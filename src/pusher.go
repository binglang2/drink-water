/**
 * ding bot
 *
 * @author binglang
 */
package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

const dingBotApi = "https://oapi.dingtalk.com/robot/send?access_token="
const defaultToken = ""

type pusher interface {
	push(content string, user *User, atAll bool) error
}

type dingBot struct {
}

func (d dingBot) push(content string, user *User, atAll bool) error {
	log.Printf("content=%s", content)
	msg := "{\"msgtype\":\"text\",\"text\":{\"content\":\"ctMsg\"}, \"at\":{\"atMobiles\":[mobile],\"isAtAll\":atAll}}"
	var token, mobile string
	if user != nil {
		token = user.DtToken
		mobile = user.Mobile
	}
	if token == "" {
		token = defaultToken
	}
	msg = strings.Replace(msg, "ctMsg", content, -1)
	msg = strings.Replace(msg, "mobile", mobile, -1)
	msg = strings.Replace(msg, "atAll", strconv.FormatBool(atAll), -1)
	if token != "" {
		_, err := http.Post(dingBotApi+token, "application/json", strings.NewReader(msg))
		if err != nil {
			return err
		}
	}
	return nil
}
