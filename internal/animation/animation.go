package animation

import tea "github.com/charmbracelet/bubbletea"

type IAnimation interface {
	Init() tea.Cmd

	Name() string

	Update(teaMsg tea.Msg) (IAnimation, tea.Cmd)

	New() IAnimation

	View() string
}
