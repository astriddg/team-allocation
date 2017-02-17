package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func addToMatches(person Person) {
	if len(matches) != 0 {
		for _, p := range people {
			var match Match
			if p != person {
				match = Match{
					Match: [2]Person{p, person},
				}
				if p.Department == person.Department {
					match.Score = 5
				} else {
					match.Score = 0
				}
				matches = append(matches, match)
			}
		}
	}
}

func delFromMatches(person Person) {
	if len(matches) != 0 {
		for k, m := range matches {
			if m.Match[0] == person || m.Match[1] == person {
				matches = append(matches[0:k], matches[k+1:]...)
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
		l := len(t.Members) - 1
		for i := 0; i < l; i++ {
			for j := 0; j < l-i; j++ {
				firstPers := getPers(t.Members[i].Name)
				secondPers := getPers(t.Members[j].Name)
				if t.Members[i].Department == t.Members[j].Department {
					firstPers.Score += 2
					secondPers.Score += 2
				} else {
					firstPers.Score++
					secondPers.Score++
				}
			}
		}
	}
}

func getDept(name string) *Department {
	for i := 0; i < len(departments); i++ {
		if departments[i].Name == name {
			return &departments[i]
		}
	}
	return &Department{}
}

func getPers(name string) *Person {
	for i := 0; i < len(people); i++ {
		if people[i].Name == name {
			return &people[i]
		}
	}
	return &Person{}
}
