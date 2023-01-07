package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tobi187/sshmanager/services"
)

type keyfiles struct {
	choices []string
	choosen int
}

func newKeyFiles() keyfiles {
	return keyfiles{
		choosen: -1,
		choices: services.FindKeyList(),
	}
}

func updateKeys(msg tea.Msg, m view_model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.keys.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.cursor = 0
			m.keys.choosen = m.cursor
			m.view_name = keyView
		}
	}
	return m, nil
}

func viewKeys(m view_model) string {
	s := "Which keyFile to use?\n\n"

	for i, choice := range m.keys.choices {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("  %s %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
