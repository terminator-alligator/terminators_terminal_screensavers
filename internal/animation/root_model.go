package animation

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
	"main.go/internal/animation/base"
)

// TODO: add the ability to chaing animations

type RootModel struct {
	Config           config.AppConfig
	frameRate        float64
	CurrentAnim      base.IAnimation
	AvailableAnim    []base.IAnimation
	CurrentAnimIndex int
	width            int
	height           int
}

// these seem like reasonable defaults
// this is done to avoid checking the size of the window in the animation
var (
	minWindowWidth  int = 20
	minWindowHeight int = 20
)

func NewRootModel(AppConfig config.AppConfig, initialAnnim base.IAnimation) RootModel {
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
	return base.TickCmd(m.frameRate)
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
	case base.TickMsg:
		var cmd tea.Cmd
		if m.height >= minWindowHeight && m.width >= minWindowWidth {
			m.CurrentAnim, cmd = m.CurrentAnim.Update(msg)
			return m, tea.Batch(base.TickCmd(m.frameRate), cmd)
		} else {
			return m, base.TickCmd(m.frameRate)
		}
	case base.AnimationFinishedMsg:
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
