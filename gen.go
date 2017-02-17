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

		orderedPeople := orderPeople()
		sort.Sort(sort.Reverse(matches))

		teamNumber := int(math.Ceil(float64(len(orderedPeople)) / float64(teamSize)))

		teams := make([]Team, teamNumber)

		for i := 0; i < teamNumber; i++ {
			selected := []Person{orderedPeople[0]}
			orderedPeople = orderedPeople[1:]
			// teamScore := 0
			for j := 1; j < teamSize; j++ {
				next, index, err := getMatchingPerson(selected, orderedPeople)
				if err != nil {
					fmt.Println(err)
				}
				if err == nil {
					selected = append(selected, next)
					orderedPeople = append(orderedPeople[:index], orderedPeople[index+1:]...)
				}

				// teamScore := next.Score
			}
			teams[i] = Team{
				Members: selected,
				// TODO: add the per-team score
			}
		}

		fmt.Println(teams)

		return nil

	}
	return fmt.Errorf("Wrong number of arguments!")

}

func getMatchingPerson(array []Person, orderedPeople []Person) (Person, int, error) {
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
	if len(leaderboard) != 0 {
		return leaderboard[0].Person, leaderboard[0].Index, nil
	} else {
		nothing := Person{}
		return nothing, 0, fmt.Errorf("No more leaders here!")
	}

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
	return false
}
