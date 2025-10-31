package animation

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO: add the ability to chaing animations

type RootModel struct {
	CurrentAnim IAnimation
	width       int
	height      int
}

// these seem like reasonable defaults
// this is done to avoid checking the size of the window in the animation
var (
	minWindowWidth  int = 20
	minWindowHeight int = 20
)

type TickMsg time.Time

// controls animations speed
// TODO: use custom speed in each animation
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type animationFinishedMsg struct{}

func AnimationFinishedCmd() tea.Cmd {
	return func() tea.Msg { return animationFinishedMsg{} }
}

func NewRootModel(initialAnnim IAnimation) RootModel {
	return RootModel{
		CurrentAnim: initialAnnim,
	}
}

func (m *RootModel) Init() tea.Cmd {
	m.CurrentAnim.Init()
	return tickCmd()
}

func (m *RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		m.width, m.height = msg.Width, msg.Height
		if m.height >= minWindowHeight && m.width >= minWindowWidth {
			m.CurrentAnim, cmd = m.CurrentAnim.Update(msg)
			return m, cmd
		} else {
			return m, nil
		}
	case TickMsg:
		var cmd tea.Cmd
		if m.height >= minWindowHeight && m.width >= minWindowWidth {
			m.CurrentAnim, cmd = m.CurrentAnim.Update(msg)
			return m, tea.Batch(tickCmd(), cmd)
		} else {
			return m, tickCmd()
		}
	case animationFinishedMsg:
		// TODO: handle animation finished, e.g. chain to next animation
		return m, nil
	}
	return m, nil
}

func (m *RootModel) View() string {
	if m.height >= minWindowHeight && m.width >= minWindowWidth {
		return m.CurrentAnim.View()
	}
	return fmt.Sprintf("Window is too small minimum width: %d, minimum height: %d", minWindowWidth, minWindowHeight)
}
