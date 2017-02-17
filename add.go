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
			// Add a new department
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
		fmt.Println("Looks like you've got the wrong arguments here")
	}
	return nil
}

func addDepartment(deptName string) (string, error) {
	if !departmentExists(deptName) {
		dept := Department{
			Name:         deptName,
			NumberPeople: 0,
		}
		departments = append(departments, dept)

		err := persistLoad()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("The department %s has been successfully added", dept.Name), nil
	} else {
		return "", fmt.Errorf("This department exists already.")
	}
}

func addPerson(persName string, deptName string) (string, error) {
	if !personExists(persName) {
		if departmentExists(deptName) {
			pers := Person{
				Name:       persName,
				Department: deptName,
				Score:      0,
			}

			people = append(people, pers)

			dept := getDept(deptName)

			dept.NumberPeople++

			addToMatches(pers)

			err := persistLoad()
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

func departmentExists(name string) bool {
	for _, dept := range departments {
		if dept.Name == name {
			return true
		}
	}
	return false
}

func personExists(name string) bool {
	for _, pers := range people {
		if pers.Name == name {
			return true
		}
	}
	return false
}
