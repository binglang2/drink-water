/**
 * @author binglang
 */
package main

import (
	"errors"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"net/http"
)

var p pusher
var users []*User

func init()  {
	p = new(dingBot)
	users, _ = SelectUserList()
}

func main() {
	registerSchedule()
	http.HandleFunc("/", index)
	http.HandleFunc("/user", addUser)
	http.HandleFunc("/name", updateName)
	http.HandleFunc("/joke", addJoke)
	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		log.Fatal("err: ", err)
	}
}

func index(respWr http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	name := req.Form.Get("name")
	defer req.Body.Close()
	result := "Hello world! "
	if name != "" {
		result += name
	}
	_, _ = respWr.Write([]byte(result))
}

func addUser(respWr http.ResponseWriter, req *http.Request)  {
	_ = req.ParseForm()
	result := "success"
	// todo: addUser
	_, _ = respWr.Write([]byte(result))
}

func updateName(respWr http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	name := req.Form.Get("name")
	mobile := req.Form.Get("m")
	result := "success"
	err := UpdateNameByMobile(name, mobile)
	if err != nil {
		result = err.Error()
	}
	_, _ = respWr.Write([]byte(result))
}

func addJoke(respWr http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	s := req.Form.Get("s")
	defer req.Body.Close()
	result := "success"
	if s != "" {
		err := AddJoke(s)
		if err != nil {
			result = err.Error()
		}
	}
	_, _ = respWr.Write([]byte(result))
}

func registerSchedule() {
	cr := cron.New()
	_, err := cr.AddFunc("2 10,12,14,16,17 ? * MON-FRI", func() {
		onTimeDrink()
	})
	if err != nil {
		panic(errors.New(err.Error()))
	}
	_, err = cr.AddFunc("22 9-17 ? * MON-FRI", func() {
		randomDrink()
	})
	if err != nil {
		panic(errors.New(err.Error()))
	}
	_, _ = cr.AddFunc("22 22 ? * MON-FRI", func() {
		SpiderJoke()
	})
	cr.Start()
}

func onTimeDrink() {
	content, _ := randomJoke()
	_ = p.push(content, nil, true)
}

func randomDrink() {
	content, _ := randomJoke()
	user, _ := randomUser()
	if user != nil {
		if user.Name == "" {
			if user.Sex == 1 {
				content = "小伙砸，" + content
			} else if user.Sex == 2 {
				content = "阿姨，" + content
			}
		} else {
			content = user.Name + "，" + content
		}
	}
	_ = p.push(content, user, false)
}

func randomJoke() (string, error) {
	jokes, err := SelectJokeList()
	if err != nil {
		log.Println("randomJoke err: e=", err)
		return "", err
	}
	if len(jokes) == 0 {
		return "nil", nil
	}
	randNum := rand.Intn(len(jokes))
	content := jokes[randNum]
	return content, nil
}

func randomUser() (*User, error) {
	if len(users) == 0 {
		users, err = SelectUserList()
		if err != nil {
			log.Println("randomUser err: e=", err)
			return nil, err
		}
	}
	randNum := rand.Intn(len(users))
	user := users[randNum]
	if len(users) > 1 {
		users = append(users[:randNum], users[randNum+1:]...)
	} else {
		users, _ = SelectUserList()
	}
	return user, nil
}
