package testanimation

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
	"main.go/internal/animation/base"
)

type TestAnimation struct {
	base.Base
}

func (m *TestAnimation) Init() tea.Cmd {
	return nil
}

// New implements the base.IAnimation interface.
func (m *TestAnimation) New(appConfig config.AppConfig) base.IAnimation {
	return &TestAnimation{
		Base: base.Base{Config: appConfig, TimeScale: appConfig.Pipes.TimeScale},
	}
}

func (m *TestAnimation) GetTimeScale() float64 {
	return m.TimeScale
}

func (m *TestAnimation) Name() string {
	return "TestanImation"
}
