package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func (s show) F(l *liner.State, fields []string) error {

	if len(fields) == 2 {
		switch fields[1] {
		case "people":
			fmt.Println(people)
			return nil
		case "departments":
			fmt.Println(departments)
			return nil
		default:
			return fmt.Errorf("That's not a listable type..")
		}

	}

	return fmt.Errorf("Wrong number of arguments!")

}
