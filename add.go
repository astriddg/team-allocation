package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func (a add) F(l *liner.State, fields []string) error {

	if fields[1] == "department" {
		if len(fields) < 3 {
			fmt.Errorf("Have you forgotten to add a department name?")
		} else if len(fields) > 3 {
			fmt.Errorf("Too many arguments!")
		} else {
			dept, err := addDepartment(fields[2])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dept)
			return nil
		}
	} else if fields[1] == "person" {
		if len(fields) < 4 {
			fmt.Errorf("Have you forgotten to add a person or department name?")
		} else if len(fields) > 4 {
			fmt.Errorf("Too many arguments!")
		} else {
			pers, err := addPerson(fields[2], fields[3])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(pers)
		}
	} else {
		fmt.Errorf("Looks like you've got the wrong arguments here")
	}
	return nil
}

func addDepartment(deptName string) (string, error) {
	if _, ok := departments[deptName]; !ok {
		dept := Department{
			Name:         deptName,
			NumberPeople: 0,
		}
		departments[deptName] = dept

		err := persistDeptAndPers()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("The department %s has been successfully added", dept.Name), nil
	} else {
		return "", fmt.Errorf("This department exists already.")
	}
}

func addPerson(persName string, deptName string) (string, error) {
	if _, ok := people[persName]; !ok {
		if dept, ok := departments[deptName]; ok {
			pers := Person{
				Name:       persName,
				Department: deptName,
				Score:      0,
			}
			people[persName] = pers

			dept.NumberPeople++

			departments[deptName] = dept

			err := persistDeptAndPers()
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("The person %s has been successfully added", pers.Name), nil

		} else {
			return "", fmt.Errorf("The department you entered does not exist.")
		}
	} else {
		return "", fmt.Errorf("This person exists already.")
	}
}
