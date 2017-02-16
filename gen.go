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
		teamSize, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		orderedDepts := orderDepts()
		orderedPeople := orderPeople()

		getOrderByDept(orderedDepts, teamSize)

		var teams []Team

		for i := 0; i < teamNumber; i++ {
			team := Team{
				Score: 0,
			}
			teams := append(teams, team)
		}

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func getOrderByDept(orderedDepts []Department, teamSize int) {
	var teamOrg []string
	var teamSlice [][]string

	for _, d := range orderedDepts {
		for i := 0; i < d.NumberPeople; i++ {
			teamOrg = append(teamOrg, d.Name)
		}
	}

	teamNumber := int(math.Ceil(float64(len(teamOrg)) / float64(teamSize)))

}

func orderDepts() []Department {
	var slice Departments
	for _, k := range departments {
		slice = append(slice, k)
	}

	sort.Sort(sort.Reverse(slice))

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
