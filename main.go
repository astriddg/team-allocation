package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"github.com/nlopes/slack"
)

// TODO: - mutex,
// - handle errors
// - generate single cid

var cmds = map[string]Command{
	"add":    add{},
	"delete": del{},
	"show":   show{},
	"gen":    gen{},
	"help":   help{},
}

type Data struct {
	People
	Departments []*Department
	Teams       []Team
	Matches     Matches
	Client      *redis.Client
	Me          string
	Locked      bool
	RTM         *slack.RTM
	CallbackId  string
	Fields      []string
}

var data = Data{}

func init() {

	connectClient()

	// Retrieve all the "tables"
	err := getList("people", data.People)
	if err != nil {
		fmt.Println(err)
	}

	err = getList("departments", data.Departments)
	if err != nil {
		fmt.Println(err)
	}

	err = getList("matches", data.Matches)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	http.HandleFunc("/", mainHandler)

	http.HandleFunc("/resp", responseHandler)

	// go handleRtm(rtm)

	// for {
	// 	time.Sleep(time.Second)
	// }

}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	api := slack.New(os.Getenv("TA_SLACK_URL"))

	list, _ := api.GetUsers()
	for _, u := range list {
		if u.Name == "team-allocation" {
			data.Me = u.ID
			break
		}
	}
	data.RTM = api.NewRTM()

	go data.RTM.ManageConnection()

	token := os.Getenv("VERIF_TOKEN")

	incomingToken := r.FormValue("token")

	if token != incomingToken {
		http.Error(w, "Wrong token", http.StatusBadRequest)
	}

	if r.FormValue("command") != "/team-allocation" {
		http.Error(w, "Wrong command", http.StatusBadRequest)
	}

	// Add a help

	if r.FormValue("text") == "" {
		http.Error(w, "No text", http.StatusBadRequest)
	}

	err := data.execute(r.FormValue("text"))

	if err != nil {
		// TODO
	}

}

func (d *Data) execute(cmd string) error {
	// Turn command string to slice
	data.Fields = strings.Fields(cmd)
	firstarg := data.Fields[0]

	if _, in := cmds[firstarg]; !in {
		return fmt.Errorf("Oops, command doesn't exist")
	}

	// retrieve command
	command := cmds[firstarg]

	// send to the right function
	err := command.Action(data.Fields)
	if err != nil {
		return err
	}

	return nil

}

// Opens the right file and reads the bytes into a struct
func getList(source string, v interface{}) error {

	value, err := data.Client.Get(source).Bytes()
	if err != nil {
		return err
	}

	err = json.Unmarshal(value, &source)
	if err != nil {
		return err
	}

	return nil

}

func connectClient() {
	data.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

	api := slack.New(os.Getenv("TA_SLACK_URL"))

	data.RTM = api.NewRTM()

	go data.RTM.ManageConnection()

	token := os.Getenv("VERIF_TOKEN")

	incomingToken := r.FormValue("token")

	if token != incomingToken {
		http.Error(w, "Wrong token", http.StatusBadRequest)
	}

	// TODO: check it's called that
	if r.FormValue("callback_id") != data.CallbackId {
		http.Error(w, "Wrong command", http.StatusBadRequest)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//TODO
	}

	response := &ButtonResponse{}

	err = json.Unmarshal(body, response)
	if err != nil {
		//TODO
	}

	if response.Actions[0].Value == "yes" {
		if data.Locked == true {
			persistTeams(data.Teams)
			data.RTM.NewOutgoingMessage("Then that's your team!", "general")
			data.Locked = false
		}
	} else if response.Actions[0].Value == "no" {
		persistTeams(data.Teams)
		cmds["gen"].Action(data.Fields)
	}

	r.Body.Close()

}
