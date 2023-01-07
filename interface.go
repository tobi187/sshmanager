package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"

// 	tea "github.com/charmbracelet/bubbletea"
// )

// // ViewNames
// const (
// 	ipView    = "IPView"
// 	keyView   = "KEYView"
// 	entryView = "NEWENTRYView"
// 	sshView   = "SSHView"
// )

// type view_model struct {
// 	choices      []string
// 	cursor       int
// 	view_name    string
// 	user_choices user_choices
// }

// type user_choices struct {
// 	server_name string
// 	key_file    string
// }

// func initialModel() view_model {
// 	return view_model{
// 		choices:   append([]string{"New Entry"}, FindIpList()...),
// 		view_name: ipView,
// 	}
// }

// func (m view_model) Init() tea.Cmd {
// 	return nil
// }

// func (m view_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	// Make sure these keys always quit
// 	if msg, ok := msg.(tea.KeyMsg); ok {
// 		k := msg.String()
// 		if k == "q" || k == "esc" || k == "ctrl+c" {
// 			// m.Quitting = true
// 			return m, tea.Quit
// 		}
// 	}

// 	switch m.view_name {
// 	case ipView, keyView:
// 		return UpdateListView(msg, m)
// 	}
// 	return m, tea.Quit
// }

// func UpdateListView(msg tea.Msg, m view_model) (tea.Model, tea.Cmd) {

// 	switch msg := msg.(type) {

// 	case tea.KeyMsg:

// 		switch msg.String() {

// 		case "up", "k":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}

// 		case "down", "j":
// 			if m.cursor < len(m.choices)-1 {
// 				m.cursor++
// 			}

// 		case "enter", " ":
// 			switch m.view_name {
// 			case ipView:
// 				if m.cursor == 0 {
// 					m.view_name = entryView
// 				} else {
// 					m.cursor = 0
// 					m.user_choices.server_name = m.choices[m.cursor]
// 					m.choices = FindKeyList()
// 					m.view_name = keyView
// 				}
// 			}
// 		}
// 	}
// 	return m, nil
// }

// func (m view_model) View() string {
// 	s := ""
// 	switch m.view_name {
// 	case keyView, ipView:
// 		s = ListView(m)
// 	}

// 	return s
// }

// func ListView(m view_model) string {
// 	s := "Which Server to connect to?\n\n"

// 	if m.view_name == keyView {
// 		s = "Which keyFile to use?\n\n"
// 	}

// 	for i, choice := range m.choices {

// 		cursor := " " // no cursor
// 		if m.cursor == i {
// 			cursor = ">" // cursor!
// 		}

// 		s += fmt.Sprintf("  %s %s\n", cursor, choice)
// 	}

// 	s += "\nPress q to quit.\n"

// 	return s
// }

// type editorFinishedMsg struct{ err error }

// func (m view_model) ssh_connection() tea.Cmd {

// 	command := exec.Command("ssh", "-i")
// 	return tea.ExecProcess(command, func(err error) tea.Msg {
// 		return editorFinishedMsg{err}
// 	})
// }

// func Gay() {
// 	p := tea.NewProgram(initialModel())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}
// }
