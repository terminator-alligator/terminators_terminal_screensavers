package matrixrain

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/animation/base"
)

func (m *MatrixRain) Update(teaMsg tea.Msg) (base.IAnimation, tea.Cmd) {
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
		if m.Config.Pipes.FrameLimit != 0 {
			if m.FrameCount >= m.Config.Pipes.FrameLimit {
				return m, base.AnimationFinishedCmd()
			}
		}
	}
	return m, nil
}
