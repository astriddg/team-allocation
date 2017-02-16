package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"

	"github.com/peterh/liner"
)

var (
	people      = map[string]Person{}
	departments = map[string]Department{}
	teams       = map[string]Team{}
	matches     = map[string]Match{}
	l           *liner.State
	cmds        = map[string]Command{
		"add":    add{},
		"delete": del{},
		"show":   show{},
		"gen":    gen{},
	}
)

var fileNames = map[string]string{
	"people":      "src/people.txt",
	"departments": "src/departments.txt",
	"teams":       "src/teams.txt",
	"matches":     "src/matches.txt",
}

func init() {
	err := getList(fileNames["people"], &people)
	if err != nil {
		fmt.Println(err)
	}

	err = getList(fileNames["departments"], &departments)
	if err != nil {
		fmt.Println(err)
	}

	err = getList(fileNames["teams"], &teams)
	if err != nil {
		fmt.Println(err)
	}

	err = getList(fileNames["matches"], &matches)
	if err != nil {
		fmt.Println(err)
	}

	l = liner.NewLiner()

	l.SetCtrlCAborts(true)

}

func main() {
	defer l.Close()

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
		cmd, err := l.Prompt("Who are you?")

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

		err = execute(l, cmd)
		if err != nil {
			fmt.Println(err)
		}

		lastExit = false

	}

}

func execute(l *liner.State, cmd string) error {
	fields := strings.Fields(cmd)
	firstarg := fields[0]

	if _, in := cmds[firstarg]; !in {
		return fmt.Errorf("Oops, command doesn't exist")
	}

	command := cmds[firstarg]

	err := command.F(l, fields)
	if err != nil {
		return err
	}

	return nil

}

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
