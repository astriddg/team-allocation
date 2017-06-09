package bot

import (
	"fmt"
	"strings"

	"github.com/team-allocation/control"
	"github.com/team-allocation/robots"
)

type bot struct{}

func init() {
	s := &bot{}
	robots.RegisterRobot("team-allocation", s)
}

func (r bot) Run(p *robots.Payload) (slashCommandImmediateReturn string) {
	text := strings.TrimSpace(p.Text)

	resp, err := control.data.execute(text, p)
	if err != nil {
		fmt.Errorf("[slackerror] %v", err)
	}

	resp.Sent()

	return "Zhboom"
}

func (r bot) Description() (description string) {
	return "The bot command is a helper bot that can invoke other bots. Useful if you are integration limited."
}
