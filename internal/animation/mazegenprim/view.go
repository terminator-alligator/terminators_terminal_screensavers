package mazegenprim

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var wallStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

var endStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(1)).
	Foreground(lipgloss.ANSIColor(0))

var startStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(4)).
	Foreground(lipgloss.ANSIColor(0))

var openStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

var (
	wallRender  = wallStyle.Render(wallChar)
	endRender   = endStyle.Render(endChar)
	startRender = startStyle.Render(startChar)
	openRender  = openStyle.Render(openChar)
)

func (m *MazeGenerationPrims) View() string {
	// Don't render if grid isn't initialized yet
	if m.Width == 0 || m.Height == 0 || len(m.grid) == 0 {
		return ""
	}

	var view strings.Builder
	// Pre-allocate with a reasonable capacity
	view.Grow(m.Width * m.Height * 12)

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.startPos[0] == x && m.startPos[1] == y {
				view.WriteString(startRender)
				continue
			}
			if m.endPos[0] == x && m.endPos[1] == y {
				view.WriteString(endRender)
				continue
			}
			switch m.grid[x][y] {
			case wall:
				view.WriteString(wallRender)
			case open:
				view.WriteString(openRender)
			}
		}
		if y < m.Height-1 {
			view.WriteRune('\n')
		}
	}
	return view.String()
}
