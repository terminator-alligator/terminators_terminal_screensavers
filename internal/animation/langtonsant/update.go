package langtonsant

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/animation/base"
)

func (m *LangtonsAnt) Update(teaMsg tea.Msg) (base.IAnimation, tea.Cmd) {
	switch msg := teaMsg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.simInit()
	case base.TickMsg:
		m.FrameCount++
		m.simUpdate()
	}
	return m, nil
}
