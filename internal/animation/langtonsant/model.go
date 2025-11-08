package langtonsant

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"main.go/config"
	"main.go/internal/animation/base"
)

type cellState int

const (
	dead cellState = iota
	alive
)

const (
	aliveChar = `█`
	deadChar  = ` `
	antChar   = `@`
)

type vec2 struct {
	x int
	y int
}

var antStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0)).
	Foreground(lipgloss.ANSIColor(9))

var cellStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

type LangtonsAnt struct {
	base.Base
	styledAlive string
	styledEmpty string
	styledAnt   string
	grid        [][]cellState
	antPos      vec2
	antDir      vec2
}

func (m *LangtonsAnt) Init() tea.Cmd {
	return nil
}

// New implements the base.IAnimation interface.
func (m *LangtonsAnt) New(appConfig config.AppConfig) base.IAnimation {
	return &LangtonsAnt{
		Base:        base.Base{Config: appConfig, TimeScale: appConfig.LangtonsAnt.TimeScale},
		styledAlive: cellStyle.Render(aliveChar),
		styledEmpty: cellStyle.Render(deadChar),
		styledAnt:   antStyle.Render(antChar),
	}
}

func (m *LangtonsAnt) GetTimeScale() float64 {
	return m.TimeScale
}

func (m *LangtonsAnt) Name() string {
	return "LangtonsAnt"
}
