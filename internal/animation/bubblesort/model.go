package bubblesort

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"main.go/internal/animation"
)

type BubbleSort struct {
	animation.Base
	grid       []int
	items      []int
	curentIndx int
}

var bgStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

func (m *BubbleSort) Init() tea.Cmd {
	return nil
}

// New implements the animation.IAnimation interface.
func (m *BubbleSort) New() animation.IAnimation {
	return &BubbleSort{}
}

func (m *BubbleSort) Name() string {
	return "BubbleSort"
}
