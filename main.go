package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"

	"github.com/peterh/liner"
)

var (
	people      = []Person{}
	departments = []Department{}
	teams       = []Team{}
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
)

// List of files to source
var fileNames = map[string]string{
	"people":      "src/people.txt",
	"departments": "src/departments.txt",
	"matches":     "src/matches.txt",
}

var words = []string{"add", "delete", "show", "gen", "person", "department", "people", "departments", "help"}

func init() {

	// Retrieve all the "tables"
	err := getList(fileNames["people"], &people)
	if err != nil {
		fmt.Println(err)
	}

	err = getList(fileNames["departments"], &departments)
	if err != nil {
		fmt.Println(err)
	}

	err = getList(fileNames["matches"], &matches)
	if err != nil {
		fmt.Println(err)
	}

	line = liner.NewLiner()

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range words {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(history_fn); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	line.SetCtrlCAborts(true)

}

func main() {
	defer line.Close()

	quitCh := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for {
			<-c
			quitCh <- true
		}
	}()

	fmt.Println("Hello you!")

	var lastExit = false
	for {
		cmd, err := line.Prompt("What do you want to do?  --  ")

		if err != nil && err.Error() == "EOF" {
			fmt.Println("bye!")
			return
		} else if err == liner.ErrPromptAborted {
			if lastExit {
				return
			} else {
				lastExit = true
				continue
			}
		} else if err != nil {
			fmt.Println("Au revoir!")
			panic(err)
		}

		if cmd == "" {
			continue
		}

		// Execute each command
		err = execute(line, cmd)
		if err != nil {
			fmt.Println(err)
		}

		if f, err := os.Create(history_fn); err != nil {
			log.Print("Error writing history file: ", err)
		} else {
			line.WriteHistory(f)
			f.Close()
		}

		lastExit = false

	}

}

func execute(line *liner.State, cmd string) error {
	// Turn command string to slice
	fields := strings.Fields(cmd)
	firstarg := fields[0]

	if _, in := cmds[firstarg]; !in {
		return fmt.Errorf("Oops, command doesn't exist")
	}

	// retrieve command
	command := cmds[firstarg]

	// send to the right function
	err := command.Action(line, fields)
	if err != nil {
		return err
	}

	return nil

}

// Opens the right file and reads the bytes into a struct
func getList(source string, v interface{}) error {
	file, err := os.OpenFile(source, os.O_CREATE|os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}

	return nil

}
