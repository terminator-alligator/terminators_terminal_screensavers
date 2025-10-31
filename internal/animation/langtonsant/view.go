package langtonsant

import "strings"

func (m *LangtonsAnt) View() string {
	// Don't render if grid isn't initialized yet
	if len(m.grid) == 0 || m.Height == 0 {
		return ""
	}

	var view strings.Builder
	// Pre-allocate with a reasonable capacity
	view.Grow(m.Width * m.Height * 4)

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.antPos.x == x && m.antPos.y == y {
				view.WriteString(m.styledAnt)
			} else {
				if m.grid[x][y] == alive {
					view.WriteString(m.styledAlive)
				} else {
					view.WriteString(m.styledEmpty)
				}
			}
		}
		if y < m.Height-1 {
			view.WriteRune('\n')
		}
	}
	return view.String()
}
