package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/peterh/liner"
)

func (g gen) Action(rtm, *slack.RTM, fields []string) error {

	if len(fields) >= 2 {
		teamSize, err := strconv.Atoi(fields[1])
		if err != nil {
			return err
		}

		var absentNames []string
		if len(fields) > 2 {
			if fields[2] == "-without" {
				absentNames = fields[3:]
			} else {
				return fmt.Errorf("Are you sure you entered the right arguments?")
			}
		}

		teams := getTeams(teamSize, absentNames)

		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Here are the teams as generated, do you like them? ")
		fmt.Println(" ")
		for _, v := range teams {
			fmt.Println(" ")
			fmt.Print("[ ")
			for _, m := range v.Members {
				fmt.Print(m.Name)
				fmt.Print("  ")
			}
			fmt.Print("] ")
			fmt.Print(v.Score)
			fmt.Println(" ")

		}
		fmt.Println(" ")
		check, err := line.Prompt("Do you like them? Shall I persist them? ")
		if err != nil {
			return fmt.Errorf("something went wrong here: %v", err)
		}
		if check == "yes" || check == "YES" || check == "Yes" || check == "yess" || check == "yes!" {
			fmt.Println("thanks!")
			persistTeams(teams)
		}

		return nil

	}
	return fmt.Errorf("Wrong number of arguments!")

}

func getTeams(teamSize int, absentees []string) []Team {
	// Get a slice of all the people in the order of the person with the highest score first
	orderedPeople := orderPeople(absentees)
	sort.Sort(sort.Reverse(matches))

	// Get the number of teams by dividing the number of people by team size.
	teamNumber := int(math.Ceil(float64(len(orderedPeople)) / float64(teamSize)))

	teams := make([]Team, teamNumber)

	// We're iterating per team then per row, such that all first lines of teams are filled
	// before the second lines are.
	for i := 0; i < teamSize; i++ {
		for j := 0; j < teamNumber; j++ {
			if i == 0 {
				// If we're in the first row, we need to create the team first, and we put in the person with the higest
				// score.
				teams[j] = Team{
					Members: []*Person{orderedPeople[0]},
				}
				orderedPeople = orderedPeople[1:]
			} else {
				// In subsequent rows, we implement the logic to get matching teammates.
				next, index, nextScore, err := getMatchingPerson(teams[j].Members, orderedPeople)
				if err != nil && err.Error() != "No more leaders here!" {
					fmt.Println(err)
				}
				if err == nil {
					teams[j].Members = append(teams[j].Members, next)
					teams[j].Score += nextScore
					orderedPeople = append(orderedPeople[:index], orderedPeople[index+1:]...)
				}
			}
		}
	}

	return teams
}

func getMatchingPerson(array []*Person, orderedPeople []*Person) (*Person, int, int, error) {
	var leaderboard Leaderboard
	for k, p := range orderedPeople {
		if personNotSelected(array, p) {
			var personTotal int
			for _, m := range matches {
				if doesMatch(m, array, p) {
					personTotal = m.Score
					if m.Match[0].Department == m.Match[0].Department {
						personTotal += 2
					} else {
						personTotal++
					}
				}
			}
			leader := &Leader{
				Person:     p,
				TotalScore: personTotal,
				Index:      k,
			}
			leaderboard = append(leaderboard, leader)
		}
	}
	sort.Sort(leaderboard)
	if len(leaderboard) != 0 {
		return leaderboard[0].Person, leaderboard[0].Index, leaderboard[0].TotalScore, nil
	} else {
		nothing := &Person{}
		return nothing, 0, 0, fmt.Errorf("No more leaders here!")
	}

}

func orderPeople(absentees []string) []*Person {
	var slice People
	for _, k := range people {
		absent := false
		for _, a := range absentees {
			if k.Name == a {
				absent = true
			}
		}
		if absent == false {
			slice = append(slice, k)
		}
	}

	sort.Sort(sort.Reverse(slice))

	return slice
}

func personNotSelected(array []*Person, p *Person) bool {
	for _, person := range array {
		if p == person {
			return false
		}
	}
	return true
}

func doesMatch(match *Match, array []*Person, p *Person) bool {
	if match.Match[0].Name == p.Name {
		for _, person := range array {
			if match.Match[1].Name == person.Name {
				return true
			}
		}
	}

	if match.Match[1].Name == p.Name {
		for _, person := range array {
			if match.Match[0].Name == person.Name {
				return true
			}
		}
	}
	return false
}
