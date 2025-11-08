package base

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
)

type AnimationFinishedMsg struct{}

type TickMsg time.Time

func TickCmd(frameRate float64) tea.Cmd {
	return tea.Tick(time.Second/time.Duration(frameRate), func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func AnimationFinishedCmd() tea.Cmd {
	return func() tea.Msg { return AnimationFinishedMsg{} }
}

type Base struct {
	Config            config.AppConfig
	Width             int
	Height            int
	TimeScale         float64
	FrameCount        int
	AnimationFinished bool
}

type IAnimation interface {
	Init() tea.Cmd

	Name() string

	GetTimeScale() float64

	Update(teaMsg tea.Msg) (IAnimation, tea.Cmd)

	New(config.AppConfig) IAnimation

	View() string
}
