package animation

import (
	"fmt"
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
	"main.go/internal/animation/base"
)

type RootModel struct {
	Config           config.AppConfig
	frameRate        float64
	timeScale        float64
	CurrentAnim      base.IAnimation
	AnimMap          map[string]base.IAnimation
	AnimNames        []string
	initialAnnim     string
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

func NewRootModel(AppConfig config.AppConfig, initialAnnim string) RootModel {
	return RootModel{
		frameRate:    AppConfig.Global.FrameRate,
		Config:       AppConfig,
		initialAnnim: initialAnnim,
	}
}

func (m *RootModel) setupAnimations() {
	m.AnimMap = make(map[string]base.IAnimation)
	for _, anim := range GetAvailableAnimations(m.Config) {
		m.AnimMap[anim.Name()] = anim
	}
	var availableNames []string
	// Filter and order animations based on the config
	for _, selectedName := range m.Config.Global.SelectedAnimations {
		if _, ok := m.AnimMap[selectedName]; ok {
			availableNames = append(availableNames, selectedName)
		}
	}
	m.AnimNames = availableNames
}

func (m *RootModel) NextAnim() {
	if m.Config.Global.Shuffle {
		randIndex := rand.Intn(len(m.AnimNames))
		nextAnim := m.AnimNames[randIndex]
		m.CurrentAnim = m.AnimMap[nextAnim]
	} else {
		m.CurrentAnim = m.AnimMap[m.AnimNames[m.CurrentAnimIndex]]
		if m.CurrentAnimIndex >= len(m.AnimNames)-1 {
			m.CurrentAnimIndex = 0
		} else {
			m.CurrentAnimIndex++
		}
	}
	// this is done to make sure the animation is reinitialised
	m.CurrentAnim = m.CurrentAnim.New(m.Config)
	m.CurrentAnim.Init()
	m.timeScale = m.CurrentAnim.GetTimeScale()
}

func (m *RootModel) Init() tea.Cmd {
	m.setupAnimations()
	if m.initialAnnim != "" {
		if anim, ok := m.AnimMap[m.initialAnnim]; ok {
			m.CurrentAnim = anim.New(m.Config)
			m.CurrentAnim.Init()
			m.timeScale = m.CurrentAnim.GetTimeScale()
		} else {
			m.NextAnim()
		}
	} else {
		m.NextAnim()
	}
	return base.TickCmd(m.frameRate * m.timeScale)
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
			return m, base.TickCmd(m.frameRate * m.timeScale)
		}
	case base.AnimationFinishedMsg:
		m.NextAnim()
		// we return a window size message here
		// to make sure that the animation is properly updated
		// otherwise, it does not know the current screen dimensions
		return m, func() tea.Msg {
			return tea.WindowSizeMsg{Width: m.width, Height: m.height}
		}
	}
	return m, nil
}

func (m *RootModel) View() string {
	if m.height >= minWindowHeight && m.width >= minWindowWidth {
		return m.CurrentAnim.View()
	}
	return fmt.Sprintf("Window is too small minimum width: %d, minimum height: %d", minWindowWidth, minWindowHeight)
}
