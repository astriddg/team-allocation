package main

import "encoding/json"

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
		var delMatch Matches
		for _, m := range matches {
			// If the first or the second person in the match is the given person.
			if m.Match[0].Name == person.Name || m.Match[1].Name == person.Name {
				delMatch = append(delMatch, m)
			}
		}

		for i := 0; i < len(matches); i++ {
			match := matches[i]
			for _, d := range delMatch {
				if match == d {
					matches = append(matches[:i], matches[i+1:]...)
					i-- // Important: decrease index
					break
				}
			}
		}

	}
}

func persistLoad() error {
	deptJson, err := json.Marshal(departments)
	if err != nil {
		return err
	}

	// Saving everything in redis
	err = client.Set("departments", deptJson, 0).Err()
	if err != nil {
		return err
	}

	peopleJson, err := json.Marshal(people)

	err = client.Set("people", peopleJson, 0).Err()
	if err != nil {
		return err
	}

	matchesJson, err := json.Marshal(matches)

	err = client.Set("matches", matchesJson, 0).Err()
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
