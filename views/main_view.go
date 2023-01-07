package views

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// ViewNames
const (
	ipView    = "IPView"
	keyView   = "KEYView"
	entryView = "NEWENTRYView"
	sshView   = "SSHView"
)

type view_model struct {
	keys      keyfiles
	serv      servers
	cursor    int
	view_name string
}

func (m view_model) Init() tea.Cmd {
	return nil
}

func initialModel() view_model {
	return view_model{
		keys:      newKeyFiles(),
		serv:      newServers(),
		cursor:    0,
		view_name: ipView,
	}
}

func (m view_model) View() string {
	s := ""
	switch m.view_name {

	case ipView:
		s = viewServers(m)
	case keyView:
		s = viewKeys(m)
	}

	return s
}

func (m view_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			// m.Quitting = true
			return m, tea.Quit
		}
	}

	switch m.view_name {

	case ipView:
		return updateServers(msg, m)
	case keyView:
		return updateKeys(msg, m)
	}

	return m, tea.Quit
}

func Start() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
