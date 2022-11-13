package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type view_model struct {
	choices   []string
	cursor    int
	selected  map[int]struct{}
	view_name string
	history   []string
}

func initialModel() view_model {
	return view_model{
		// Our to-do list is a grocery list
		choices:   append(FindIpList(), "New Entry"),
		view_name: "ip",
		history:   []string{},
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m view_model) Init() tea.Cmd {
	return nil
}

func (m view_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.cursor = 0
			m.history = append(m.history, m.choices[m.cursor])
			m.choices = FindKeyList()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m view_model) View() string {
	// The header
	s := "Which Server to connect to?\n\n"

	if len(m.history) > 0 {
		s = "Which keyFile to use?\n\n"
	}

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func Gay() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
