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
