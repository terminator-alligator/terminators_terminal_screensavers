package pipes

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
	"main.go/internal/animation"
)

/*
This code is not grate.
I have been drowning my brain in tea and coffee,
so this is the best thing it could come up with

hopefully it should be fixed in the future :)
*/

type vac2 struct {
	x int
	y int
}

type Pipes struct {
	animation.Base
	grid            [][]cellState
	pipeList        []*pipe
	changDirChance  float64
	pipeSpawnChance float64
}

type pipe struct {
	moveLength int
	color      int
	pos        vac2
	dir        vac2
	prevDir    vac2
}

type dirKey struct {
	in  vac2
	out vac2
}

type cellState struct {
	dir   dirKey
	color int
}

func (m *Pipes) Init() tea.Cmd {
	m.setupColours()
	return nil
}

// New implements the animation.IAnimation interface.
func (m *Pipes) New(appConfig config.AppConfig) animation.IAnimation {
	return &Pipes{Base: animation.Base{Config: appConfig, TimeScale: appConfig.Pipes.TimeScale}}
}

func (m *Pipes) GetTimeScale() float64 {
	return m.TimeScale
}

func (m *Pipes) Name() string {
	return "Pipes"
}
