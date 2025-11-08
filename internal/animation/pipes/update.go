package pipes

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/animation/base"
)

func (m *Pipes) Update(teaMsg tea.Msg) (base.IAnimation, tea.Cmd) {
	switch msg := teaMsg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.simInit()
	case base.TickMsg:
		m.FrameCount++
		m.simUpdate()
		if m.AnimationFinished {
			return m, base.AnimationFinishedCmd()
		}
	}
	return m, nil
}
