package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/go-redis/redis"
	"github.com/nlopes/slack"
	"github.com/peterh/liner"
)

// TODO: - mutex,
// - message button,
//

var (
	people      = People{}
	departments = []*Department{}
	teams       = []*Team{}
	matches     Matches
	line        *liner.State
	cmds        = map[string]Command{
		"add":    add{},
		"delete": del{},
		"show":   show{},
		"gen":    gen{},
		"help":   help{},
	}
	history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
	client     *redis.Client
	Me         string
	mutex      = sync.Mutex{}
)

// List of files to source
var fileNames = map[string]string{
	"people":      "src/people.txt",
	"departments": "src/departments.txt",
	"matches":     "src/matches.txt",
}

var words = []string{"add", "delete", "show", "gen", "person", "department", "people", "departments", "help"}

func init() {

	connectClient()

	// Retrieve all the "tables"
	err := getList("people", &people)
	if err != nil {
		fmt.Println(err)
	}

	err = getList("departments", &departments)
	if err != nil {
		fmt.Println(err)
	}

	err = getList("matches", &matches)
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
			Me = u.ID
			break
		}
	}

	rtm := api.NewRTM()

	go rtm.ManageConnection()

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

	err := execute(rtm, r.FormValue("text"))

	if err != nil {
		//TO DO
	}

}

func execute(rtm *slack.RTM, cmd string) error {
	// Turn command string to slice
	fields := strings.Fields(cmd)
	firstarg := fields[0]

	if _, in := cmds[firstarg]; !in {
		return fmt.Errorf("Oops, command doesn't exist")
	}

	// retrieve command
	command := cmds[firstarg]

	// send to the right function
	err := command.Action(rtm, fields)
	if err != nil {
		return err
	}

	return nil

}

// Opens the right file and reads the bytes into a struct
func getList(source string, v interface{}) error {

	value, err := client.Get(source).Bytes()
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
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
