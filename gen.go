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

		stringTeams, teamNumber := getOrderByDept(orderedDepts, teamSize)

		teams := make([]Team, teamNumber)

		for i := 0; i < len(stringTeams); i++ {
            firstMember := orderedPeople[0]
            secondMember := findMatch(firstMember)
			team := Team{
                Members: 
            }
		}

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func findMatch(person Person) {
    
}

func getOrderByDept(orderedDepts []Department, teamSize int) ([][]string, int) {
	var teamOrg []string

	for _, d := range orderedDepts {
		for i := 0; i < d.NumberPeople; i++ {
			teamOrg = append(teamOrg, d.Name)
		}
	}

	teamNumber := int(math.Ceil(float64(len(teamOrg)) / float64(teamSize)))

	teamSlice := make([][]string, teamNumber)
	for i := range teamSlice {
		teamSlice[i] = make([]string, teamSize)
	}

	var j = 0
	for k := 0; k < teamSize; k++ {
		for i := 0; i < teamNumber; i++ {
			if j < len(teamOrg) {
				teamSlice[i][k] = teamOrg[j]
				j++
			}
		}
	}

	fmt.Println(teamSlice)

	return teamSlice, teamNumber

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

	sort.Sort(sort.Reverse(slice))

	return slice
}
