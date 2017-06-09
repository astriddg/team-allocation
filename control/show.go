package control

import (
	"fmt"
	"sort"

	"github.com/team-allocation/robots"
)

func (s show) Action(fields []string, p *robots.Payload) (*robots.IncomingWebhook, error) {

	if len(fields) == 2 {
		switch fields[1] {
		case "people":
			showPeople()
			return nil
		case "departments":
			showDepartment()
			return nil
		case "matches":
			showMatches()
			return nil
		default:
			return fmt.Errorf("That's not a listable type..")
		}

	}

	return fmt.Errorf("Wrong number of arguments!")

}

func showMatches() {
	sort.Sort(sort.Reverse(data.Matches))
	showString := " \n"
	for _, m := range data.Matches {
		showString += fmt.Sprintf("%s - %s : %v \n", m.Match[0].Name, m.Match[1].Name, m.Score)
	}

	data.RTM.NewOutgoingMessage(showString, "general")

}

func showPeople() {
	showString := " \n"
	for _, p := range data.People {
		showString += fmt.Sprintf("%s: %v \n", p.Name, p.Score)
	}

	data.RTM.NewOutgoingMessage(showString, "general")

}

func showDepartment() {
	showString := " \n"
	for _, d := range data.Departments {
		showString += fmt.Sprintf("%s: %v \n", d.Name, d.NumberPeople)

		fmt.Printf("%v: %v", d.Name, d.NumberPeople)
		fmt.Println(" ")
	}

	data.RTM.NewOutgoingMessage(showString, "general")
}
