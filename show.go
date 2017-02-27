package main

import (
	"fmt"
	"sort"

	"github.com/peterh/liner"
)

func (s show) Action(line *liner.State, fields []string) error {

	if len(fields) == 2 {
		switch fields[1] {
		case "people":
			fmt.Println(people)
			return nil
		case "departments":
			fmt.Println(departments)
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
	copy := matches
	sort.Sort(sort.Reverse(copy))
	fmt.Println(" ")
	for _, m := range copy {
		fmt.Print(m.Match[0].Name)
		fmt.Print(" ")
		fmt.Print(m.Match[1].Name)
		fmt.Print(" ")
		fmt.Print(m.Score)
		fmt.Println(" ")

	}
}
