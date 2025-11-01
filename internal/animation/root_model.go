package animation

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
)

// TODO: add the ability to chaing animations

type RootModel struct {
	Config      config.AppConfig
	frameRate   float64
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
func tickCmd(frameRate float64) tea.Cmd {
	return tea.Tick(time.Second/time.Duration(frameRate), func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type animationFinishedMsg struct{}

func AnimationFinishedCmd() tea.Cmd {
	return func() tea.Msg { return animationFinishedMsg{} }
}

func NewRootModel(AppConfig config.AppConfig, initialAnnim IAnimation) RootModel {
	return RootModel{
		frameRate:   AppConfig.Global.FrameRate,
		Config:      AppConfig,
		CurrentAnim: initialAnnim,
	}
}

func (m *RootModel) Init() tea.Cmd {
	m.CurrentAnim.Init()
	timeScale := m.CurrentAnim.GetTimeScale()
	if timeScale <= 0 {
		return nil
	}
	m.frameRate = m.frameRate * timeScale
	return tickCmd(m.frameRate)
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
			return m, tea.Batch(tickCmd(m.frameRate), cmd)
		} else {
			return m, tickCmd(m.frameRate)
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
