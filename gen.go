package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/peterh/liner"
)

func (g gen) F(l *liner.State, fields []string) error {

	if len(fields) == 2 {
		teamSiZe, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		TotalNbPeople := getNumberPeople()

		teamNumber := int64(math.Ceil(float64(TotalNbPeople) / float64(teamSiZe)))

		orderedDepts := orderDepts()
		orderedPeople := orderPeople

		var teams []Team

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func getNumberPeople() int {
	var totalNumber int

	for _, k := range departments {
		totalNumber += k.NumberPeople
	}

	return totalNumber
}

func orderDepts() []Department {
	var slice Departments
	for _, k := range departments {
		slice = append(slice, k)
	}

	sort.Sort(slice)

	return slice
}

func orderPeople() []Person {
	var slice People
	for _, k := range people {
		slice = append(slice, k)
	}

	sort.Sort(slice)

	return slice
}
