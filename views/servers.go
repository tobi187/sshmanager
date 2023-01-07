package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tobi187/sshmanager/services"
)

type servers struct {
	choices []string
	choosen int
}

func newServers() servers {
	return servers{
		choosen: -1,
		choices: services.FindIpList(),
	}
}

func updateServers(msg tea.Msg, m view_model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.serv.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			if m.cursor == 0 { // new entry
				m.cursor = 0
				m.view_name = entryView
			} else {
				m.keys.choosen = m.cursor
				m.view_name = keyView
				m.cursor = 0
			}
		}
	}
	return m, nil
}

func viewServers(m view_model) string {
	s := "Which Server to connect to?\n\n"

	for i, choice := range m.serv.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("  %s %s\n", cursor, choice)
	}
	s += "\nPress q to quit.\n"

	return s
}
