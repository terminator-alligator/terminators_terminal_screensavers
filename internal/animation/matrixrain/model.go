package matrixrain

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
	"main.go/internal/animation/base"
)

type MatrixRain struct {
	base.Base
	grid        [][]int
	brightness  [][]int
	trailLength int
	spawnChance float64
}

var availableChr = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&_*+-=(){}[];:'\|/<>,.?~`

func (m *MatrixRain) Init() tea.Cmd {
	m.setUpChrMap()
	return nil
}

// New implements the base.IAnimation interface.
func (m *MatrixRain) New(appConfig config.AppConfig) base.IAnimation {
	return &MatrixRain{
		Base:        base.Base{Config: appConfig, TimeScale: appConfig.MatrixRain.TimeScale},
		trailLength: appConfig.MatrixRain.TrailLength,
		spawnChance: appConfig.MatrixRain.SpawnChance,
	}
}

func (m *MatrixRain) GetTimeScale() float64 {
	return m.TimeScale
}

func (m *MatrixRain) Name() string {
	return "MatrixRain"
}
