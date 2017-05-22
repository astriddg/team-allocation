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
			showPeople()
			return nil
		case "departments":
			showDepartment()
			return nil
		case "matches":
			showMatches()
			return nil
		default:
			return fmt.Errorf("That's not a listable type..")
		}

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func showMatches() {
	sort.Sort(sort.Reverse(matches))
	fmt.Println(" ")
	for _, m := range matches {
		fmt.Print(m.Match[0].Name)
		fmt.Print(" ")
		fmt.Print(m.Match[1].Name)
		fmt.Print(" ")
		fmt.Print(m.Score)
		fmt.Println(" ")

	}
}

func showPeople() {
	fmt.Println(" ")
	for _, p := range people {
		fmt.Printf("%v: %v", p.Name, p.Score)
		fmt.Println(" ")
	}
}

func showDepartment() {
	fmt.Println(" ")
	for _, d := range departments {
		fmt.Printf("%v: %v", d.Name, d.NumberPeople)
		fmt.Println(" ")
	}
}
