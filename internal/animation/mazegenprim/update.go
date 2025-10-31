package mazegenprim

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/animation"
)

func (m *MazeGenerationPrims) Update(teaMsg tea.Msg) (animation.IAnimation, tea.Cmd) {
	switch msg := teaMsg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.simInit()
	case animation.TickMsg:
		m.FrameCount++
		m.simUpdate()
		if m.AnimationFinished {
			return m, animation.AnimationFinishedCmd()
		}
	}
	return m, nil
}
