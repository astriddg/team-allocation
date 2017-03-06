package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func (a add) Action(line *liner.State, fields []string) error {

	if fields[1] == "department" || fields[1] == "d" {

		if len(fields) < 3 {
			fmt.Errorf("Have you forgotten to add a department name?")

		} else if len(fields) > 3 {
			fmt.Errorf("Too many arguments!")

		} else {
			// Add a new department
			message, err := addDepartment(fields[2])
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(message)

			return nil
		}

	} else if fields[1] == "person" || fields[1] == "p" {

		if len(fields) < 4 {
			fmt.Errorf("Have you forgotten to add a person or department name?")

		} else if len(fields) > 4 {
			fmt.Errorf("Too many arguments!")

		} else {
			// Add a new person
			message, err := addPerson(fields[2], fields[3])
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(message)

		}

	} else {
		fmt.Println("Looks like you've got the wrong arguments here")
	}
	return nil
}

func addDepartment(deptName string) (string, error) {

	if _, ok := departmentExists(deptName); !ok {
		// Create department
		dept := &Department{
			Name:         deptName,
			NumberPeople: 0,
		}
		// Add to the list
		departments = append(departments, dept)

		// then persist the whole thing
		err := persistLoad()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("The department %s has been added successfully", dept.Name), nil
	} else {
		return "", fmt.Errorf("This department exists already.")
	}
}

func addPerson(persName string, deptName string) (string, error) {
	if _, ok := personExists(persName); !ok {
		if dept, ok := departmentExists(deptName); ok {
			pers := &Person{
				Name:       persName,
				Department: dept,
				Score:      0,
			}

			dept.NumberPeople++

			// You want a "match" to be created between this new person and all of the others.
			addToMatches(pers)
			people = append(people, pers)

			// fmt.Println(&pers)
			// Save it all
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

func departmentExists(name string) (*Department, bool) {
	for _, dept := range departments {
		if dept.Name == name {
			return dept, true
		}
	}
	return &Department{}, false
}

func personExists(name string) (*Person, bool) {
	for _, pers := range people {
		if pers.Name == name {
			return pers, true
		}
	}
	return &Person{}, false
}
