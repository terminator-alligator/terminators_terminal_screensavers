package animation

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
)

type IAnimation interface {
	Init() tea.Cmd

	Name() string

	GetTimeScale() float64

	Update(teaMsg tea.Msg) (IAnimation, tea.Cmd)

	New(config.AppConfig) IAnimation

	View() string
}
