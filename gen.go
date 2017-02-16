package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/peterh/liner"
	"github.com/y0ssar1an/q"
)

func (g gen) F(l *liner.State, fields []string) error {

	if len(fields) == 2 {
		teamSize, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		orderedDepts := orderDepts()
		// orderedPeople := orderPeople()

		getOrderByDept(orderedDepts, teamSize)

		// var teams []Team

		// for i := 0; i < teamNumber; i++ {
		// 	team := Team{
		// 		Score: 0,
		// 	}
		// 	teams := append(teams, team)
		// }

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func getOrderByDept(orderedDepts []Department, teamSize int) {
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

	q.Q(len(teamOrg))
	q.Q(teamSize)
	q.Q(teamNumber)

	var j = 0
	for k := 0; k < teamSize; k++ {
		for i := 0; i < teamNumber; i++ {
			if j < len(teamOrg) {
				q.Q(k, i, j)
				teamSlice[i][k] = teamOrg[j]
				j++
			}
		}
	}

	fmt.Println(teamSlice)

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
