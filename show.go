package main

import (
	"fmt"
	"sort"

	"github.com/nlopes/slack"
)

func (s show) Action(rtm *slack.RTM, fields []string) error {

	if len(fields) == 2 {
		switch fields[1] {
		case "people":
			showPeople(rtm)
			return nil
		case "departments":
			showDepartment(rtm)
			return nil
		case "matches":
			showMatches(rtm)
			return nil
		default:
			return fmt.Errorf("That's not a listable type..")
		}

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func showMatches(rtm *slack.RTM) {
	sort.Sort(sort.Reverse(matches))
	showString := " \n"
	for _, m := range matches {
		showString += fmt.Sprintf("%s - %s : %v \n", m.Match[0].Name, m.Match[1].Name, m.Score)
	}

	rtm.NewOutgoingMessage(showString, "general")

}

func showPeople(rtm *slack.RTM) {
	showString := " \n"
	for _, p := range people {
		showString += fmt.Sprintf("%s: %v \n", p.Name, p.Score)
	}

	rtm.NewOutgoingMessage(showString, "general")

}

func showDepartment(rtm *slack.RTM) {
	showString := " \n"
	for _, d := range departments {
		showString += fmt.Sprintf("%s: %v \n", d.Name, d.NumberPeople)

		fmt.Printf("%v: %v", d.Name, d.NumberPeople)
		fmt.Println(" ")
	}

	rtm.NewOutgoingMessage(showString, "general")
}
