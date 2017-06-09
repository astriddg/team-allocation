package control

import (
	"fmt"

	"github.com/team-allocation/robots"
)

func (d del) Action(fields []string, p *robots.Payload) (*robots.IncomingWebhook, error) {

	response := robots.IncomingWebhook{
		Domain:      p.TeamDomain,
		Channel:     p.ChannelID,
		Username:    "team-allocation",
		UnfurlLinks: true,
		Parse:       robots.ParseStyleFull,
		Text:        "Done",
	}

	if fields[1] == "department" || fields[1] == "d" {
		if len(fields) < 3 {
			return nil, fmt.Errorf("Have you forgotten to add a department name?")
		} else if len(fields) > 3 {
			return nil, fmt.Errorf("Too many arguments!")
		} else {
			// Delete department
			message, err := delDepartment(fields[2])
			if err != nil {
				return nil, err
			}

			return nil, response
		}
	} else if fields[1] == "person" || fields[1] == "p" {
		if len(fields) < 3 {
			return nil, fmt.Errorf("Have you forgotten to add a person name?")
		} else if len(fields) > 3 {
			return nil, fmt.Errorf("Too many arguments!")
		} else {
			// Delete person
			message, err := delPerson(fields[2], false)
			if err != nil {
				return nil, err
			}

			return nil, response

		}
	}
	return nil, fmt.Errorf("Looks like you've got the wrong arguments here")

}

func delDepartment(deptName string) (string, error) {
	if dept, ok := departmentExists(deptName); ok {

		// deleting first anyone from that department
		for _, v := range data.People {
			if v.Department == dept {
				// true: the department will be deleted
				delPerson(v.Name, true)
			}
		}

		// then deleting department
		for k, d := range data.Departments {
			if d.Name == deptName {
				data.Departments = append(data.Departments[:k], data.Departments[k+1:]...)
			}
		}

		err := persistLoad()
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("The department %s has been successfully deleted", deptName), nil

	} else {
		return "", fmt.Errorf("This department does not exist.")
	}
}

func delPerson(persName string, deleteDept bool) (string, error) {
	if _, ok := personExists(persName); ok {

		for k, p := range data.People {
			if p.Name == persName {
				// Reduce the number of people in that department
				if !deleteDept {
					p.Department.NumberPeople--
				}

				data.People = append(data.People[:k], data.People[k+1:]...)
				// The person also has to be deleted from the matches list.
				delFromMatches(p)

			}
		}

		err := persistLoad()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("The person %s has been deleted", persName), nil

	}
	return "", fmt.Errorf("This person does not exist.")
}
