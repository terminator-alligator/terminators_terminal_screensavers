package bubblesort

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"main.go/config"
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
func (m *BubbleSort) New(appConfig config.AppConfig) animation.IAnimation {
	return &BubbleSort{Base: animation.Base{Config: appConfig, TimeScale: appConfig.BobbleSort.TimeScale}}
}

func (m *BubbleSort) GetTimeScale() float64 {
	return m.TimeScale
}

func (m *BubbleSort) Name() string {
	return "BubbleSort"
}
