package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func (d del) Action(line *liner.State, fields []string) error {

	if fields[1] == "department" {
		if len(fields) < 3 {
			fmt.Errorf("Have you forgotten to add a department name?")
		} else if len(fields) > 3 {
			fmt.Errorf("Too many arguments!")
		} else {
			// Delete department
			_, err := delDepartment(fields[2])
			if err != nil {
				fmt.Println(err)
			}
			return nil
		}
	} else if fields[1] == "person" {
		if len(fields) < 3 {
			fmt.Errorf("Have you forgotten to add a person name?")
		} else if len(fields) > 3 {
			fmt.Errorf("Too many arguments!")
		} else {
			// Delete person
			_, err := delPerson(fields[2], false)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		return fmt.Errorf("Looks like you've got the wrong arguments here")
	}
	return nil
}

func delDepartment(deptName string) (string, error) {
	if dept, ok := departmentExists(deptName); ok {

		// deleting first anyone from that department
		for _, v := range people {
			if v.Department == dept {
				// true: the department will be deleted
				delPerson(v.Name, true)
			}
		}

		// then deleting department
		for k, d := range departments {
			if d.Name == deptName {
				departments = append(departments[:k], departments[k+1:]...)
			}
		}

		err := persistLoad()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("The department %s has been successfully deleted", deptName), nil

	} else {
		return "", fmt.Errorf("This department does not exist.")
	}
}

func delPerson(persName string, deleteDept bool) (string, error) {
	if _, ok := personExists(persName); ok {

		for k, p := range people {
			if p.Name == persName {
				// Reduce the number of people in that department
				if !deleteDept {
					p.Department.NumberPeople--
				}

				people = append(people[:k], people[k+1:]...)
				// The person also has to be deleted from the matches list.
				delFromMatches(p)

			}
		}

		err := persistLoad()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("The person %s has been deleted", persName), nil

	}
	return "", fmt.Errorf("This person does not exist.")
}
