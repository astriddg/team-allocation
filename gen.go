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
		sort.Sort(sort.Reverse(matches))

		stringTeams, teamNumber := getTeamNumber(orderedDepts, teamSize)

		teams := make([]Team, teamNumber)

		for i := 0; i < len(stringTeams); i++ {
			selected := []Person{orderedPeople[0]}
			orderedPeople = orderedPeople[1:]
			// teamScore := 0
			for j := 1; j < teamSize; j++ {
				next, int := getMatchingPerson(selected, orderedPeople)
				selected = append(selected, next)

				// teamScore := next.Score
			}
			teams[i] = Team{
				Members: selected,
				// TODO: add the per-team score
			}
		}

		fmt.Println(teams)

	} else {
		return fmt.Errorf("Wrong number of arguments!")
	}
}

func getMatchingPerson(array []Person, orderedPeople []Person) (Person, int) {
	var leaderboard Leaderboard
	for k, p := range orderedPeople {
		if personNotSelected(array, p) {
			var personTotal int
			for _, m := range matches {
				if doesMatch(m, array, p) {
					personTotal += m.Score
				}
			}
			leader := Leader{
				Person:     p,
				TotalScore: personTotal,
				Index:      k,
			}
			leaderboard = append(leaderboard, leader)
		}
	}
	sort.Sort(sort.Reverse(leaderboard))
	if len(leaderboard != 0) {
		return leaderboard[0].Person, leaderboard[0].Index
	} else {
		return nil, nil
	}

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

func personNotSelected(array []Person, p Person) bool {
	for _, person := range array {
		if p == person {
			return false
		}
	}
	return true
}

func doesMatch(match Match, array []Person, p Person) bool {
	if match.Match[0] == p {
		for _, person := range array {
			if match.Match[1] == person {
				return true
			}
		}
	}

	if match.Match[1] == p {
		for _, person := range array {
			if match.Match[0] == person {
				return true
			}
		}
	}
}
