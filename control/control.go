package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"github.com/nlopes/slack"
	"github.com/team-allocation/robots"
)

// TODO: - mutex,
// - handle errors

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

func (d *Data) execute(cmd string, p *robots.Payload)  *robots.IncomingWebhook error {
	// Turn command string to slice
	data.Fields = strings.Fields(cmd)
	firstarg := data.Fields[0]

	if _, in := cmds[firstarg]; !in {
		return nil, fmt.Errorf("Oops, command doesn't exist")
	}

	// retrieve command
	command := cmds[firstarg]

	// send to the right function
	resp, err := command.Action(data.Fields, p)
	if err != nil {
		return nil, err
	}

	return resp, nil

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
	redisUrl, err := url.Parse(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Panicf("Error parsing REDIS_URL: %v", err)
	}
	password, _ := redisUrl.User.Password()

	data.Client = redis.NewClient(&redis.Options{
		Addr:     redisUrl.Host,
		Password: password,
		DB:       0,
	})
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

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
