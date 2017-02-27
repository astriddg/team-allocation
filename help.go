package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func (h help) Action(line *liner.State, fields []string) error {

	helpMessage :=
		` 
        |---------------------------------------------- HELP ----------------------------------------------|
        |Commands:                                                                                         |
        |                                                                                                  |
        |    - show [people / departments / matches] : shows the chosen list with each individual score,   |
        |                                                                                                  |
        |    - gen [integer] : e.g. gen 4 - will generate teams of at most four.                           |
        |                    If, say there are 9 members, teams of 3 will be generated.                    |
        |                                                                                                  |
        |    - add department [departmentName] : adds a new department with no members.                    |
        |                                                                                                  |
        |    - add person [personName] [department Name] : adds a new person to the given department.      |
        |                                                                                                  |
        |    - delete [departmentName/personName] : deletes a department or a person.                      |
        |--------------------------------------------------------------------------------------------------|
        `

	fmt.Println(helpMessage)

	return nil

}
