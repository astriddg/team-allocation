package control

import "encoding/json"

type ButtonResponse struct {
	CallbackId string   `json:"callback_id"`
	Token      string   `json:"token"`
	Actions    []Action `json"actions"`
}

type Action struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func addToMatches(person *Person) {
	for i, _ := range data.People {
		p := data.People[i]
		match := &Match{
			Match: [2]*Person{p, person},
		}
		if p.Department == person.Department {
			match.Score = 5
		} else {
			match.Score = 0
		}
		data.People[i].Score = p.Score + match.Score
		person.Score = person.Score + match.Score

		data.Matches = append(data.Matches, match)

	}
}

func delFromMatches(person *Person) {
	if len(data.Matches) != 0 {
		var delMatch Matches
		for _, m := range data.Matches {
			// If the first or the second person in the match is the given person.
			if m.Match[0].Name == person.Name || m.Match[1].Name == person.Name {
				delMatch = append(delMatch, m)
			}
		}

		for i := 0; i < len(data.Matches); i++ {
			match := data.Matches[i]
			for _, d := range delMatch {
				if match == d {
					data.Matches = append(data.Matches[:i], data.Matches[i+1:]...)
					i-- // Important: decrease index
					break
				}
			}
		}

	}
}

func persistLoad() error {
	deptJson, err := json.Marshal(data.Departments)
	if err != nil {
		return err
	}

	// Saving everything in redis
	err = data.Client.Set("departments", deptJson, 0).Err()
	if err != nil {
		return err
	}

	peopleJson, err := json.Marshal(data.People)

	err = data.Client.Set("people", peopleJson, 0).Err()
	if err != nil {
		return err
	}

	matchesJson, err := json.Marshal(data.Matches)

	err = data.Client.Set("matches", matchesJson, 0).Err()
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
	for i := 0; i < len(data.Matches); i++ {
		if (data.Matches[i].Match[0].Name == firstName && data.Matches[i].Match[1].Name == secondName) ||
			(data.Matches[i].Match[1].Name == firstName && data.Matches[i].Match[0].Name == secondName) {
			return data.Matches[i]
		}
	}
	return &Match{}
}
