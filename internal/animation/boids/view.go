package boids

import (
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var bg = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

var boidsStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0)).
	Foreground(lipgloss.ANSIColor(7))

var directions = [8]string{
	boidsStyle.Render("⭢"), // Right
	boidsStyle.Render("⭧"), // Right
	boidsStyle.Render("⭡"), // Right
	boidsStyle.Render("⭦"), // Right
	boidsStyle.Render("⭠"), // Right
	boidsStyle.Render("⭩"), // Right
	boidsStyle.Render("⭣"), // Right
	boidsStyle.Render("⭨"), // Right
}
var bgRender = bg.Render(" ")

func (m *Boids) View() string {
	// Don't render if grid isn't initialized yet
	if m.Width == 0 || m.Height == 0 || len(m.grid) == 0 {
		return ""
	}

	var view strings.Builder
	// Pre-allocate with a reasonable capacity
	view.Grow(m.Width * m.Height * 12)

	// Draw boids on the grid
	for _, b := range m.flock {
		x, y := int(b.pos.x), int(b.pos.y)
		// find the closest angle
		angle := math.Atan2(-b.dir.y, b.dir.x) // Invert Y for terminal rendering
		if angle < 0 {
			angle += 2 * math.Pi
		}
		index := int(math.Round(angle/(math.Pi/4))) % 8
		m.grid[x][y] = index + 1 // Use index+1 to avoid 0
	}

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.grid[x][y] > 0 {
				view.WriteString(directions[m.grid[x][y]-1])
			} else {
				view.WriteString(bgRender)
			}
		}
		if y < m.Height-1 {
			view.WriteRune('\n')
		}
	}
	return view.String()
}
