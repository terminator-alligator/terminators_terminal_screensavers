package mazegenprim

import (
	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/animation"
)

// TODO: loock at adding a maze solving algorithm

const (
	wallChar  = ` `
	openChar  = `█`
	startChar = `#`
	endChar   = "%"
)

type gridState int

const (
	open gridState = iota
	wall
	start
	end
)

type MazeGenerationPrims struct {
	animation.Base
	grid               [][]gridState
	wallsList          [][2]int
	startPos           [2]int
	endPos             [2]int
	finishedGenerating bool
}

func (m *MazeGenerationPrims) Init() tea.Cmd {
	return nil
}

// New implements the animation.IAnimation interface.
func (m *MazeGenerationPrims) New() animation.IAnimation {
	return &MazeGenerationPrims{}
}

func (m *MazeGenerationPrims) Name() string {
	return "MazeGenerationPrims"
}
