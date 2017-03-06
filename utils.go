package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func addToMatches(person *Person) {
	for i, _ := range people {
		p := people[i]
		match := &Match{
			Match: [2]*Person{p, person},
		}
		if p.Department == person.Department {
			match.Score = 5
		} else {
			match.Score = 0
		}
		people[i].Score = p.Score + match.Score
		person.Score = person.Score + match.Score

		matches = append(matches, match)

	}
}

func delFromMatches(person *Person) {
	if len(matches) != 0 {
		for k, m := range matches {
			// If the first or the second person in the match is the given person.
			if m.Match[0].Name == person.Name || m.Match[1].Name == person.Name {
				if k < len(matches) {
					matches = append(matches[0:k], matches[k+1:]...)
					// That's if the match is at the end.
				} else {
					matches = matches[0:k]
				}
			}
		}
	}
}

func persistLoad() error {
	deptFile, err := json.Marshal(departments)

	// Saving everything in files
	err = ioutil.WriteFile(fileNames["departments"], deptFile, os.ModeType)
	if err != nil {
		return err
	}

	persFile, err := json.Marshal(people)

	err = ioutil.WriteFile(fileNames["people"], persFile, os.ModeType)
	if err != nil {
		return err
	}

	matchFile, err := json.Marshal(matches)

	err = ioutil.WriteFile(fileNames["matches"], matchFile, os.ModeType)
	if err != nil {
		return err
	}

	return nil
}

func persistTeams(teams []Team) {
	for _, t := range teams {
		l := len(t.Members)
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				firstPers := t.Members[i]
				secondPers := t.Members[j]
				match := getMatch(firstPers.Name, secondPers.Name)
				if firstPers.Department == secondPers.Department {
					firstPers.Score += 2
					secondPers.Score += 2
					match.Score += 2
				} else {
					firstPers.Score++
					secondPers.Score++
					match.Score++
				}

			}
		}
	}

	persistLoad()
}

func getMatch(firstName string, secondName string) *Match {
	for i := 0; i < len(matches); i++ {
		if (matches[i].Match[0].Name == firstName && matches[i].Match[1].Name == secondName) ||
			(matches[i].Match[1].Name == firstName && matches[i].Match[0].Name == secondName) {
			return matches[i]
		}
	}
	return &Match{}
}
