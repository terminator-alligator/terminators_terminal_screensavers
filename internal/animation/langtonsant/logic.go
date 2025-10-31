package langtonsant

// initialise the sim
func (m *LangtonsAnt) simInit() {
	if m.Width <= 5 || m.Height <= 5 {
		return
	}
	m.grid = make([][]cellState, m.Width)
	for i := range m.grid {
		m.grid[i] = make([]cellState, m.Height)
	}
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			m.grid[x][y] = dead
		}
	}
	m.antPos = vec2{x: m.Width / 2, y: m.Height / 2}
	m.antDir = vec2{x: 0, y: -1}
}

func (m *LangtonsAnt) simUpdate() {
	antNextPos := vec2{
		m.antPos.x + m.antDir.x,
		m.antPos.y + m.antDir.y,
	}

	// If ant is about to go off-screen, reset its position to the center.
	if antNextPos.x >= m.Width || antNextPos.x < 0 || antNextPos.y >= m.Height || antNextPos.y < 0 {
		m.antPos = vec2{x: m.Width / 2, y: m.Height / 2}
		return
	}

	if m.grid[antNextPos.x][antNextPos.y] == alive { // On a "live" (true) cell
		// Turn 90 degrees left (counter-clockwise)
		m.antDir.x, m.antDir.y = m.antDir.y, -m.antDir.x
	} else { // On a "dead" (false) cell
		// Turn 90 degrees right (clockwise)
		m.antDir.x, m.antDir.y = -m.antDir.y, m.antDir.x
	}

	// Flip the color of the cell the ant is leaving
	if m.grid[m.antPos.x][m.antPos.y] == alive {
		m.grid[m.antPos.x][m.antPos.y] = dead
	} else {
		m.grid[m.antPos.x][m.antPos.y] = alive
	}
	// Move ant to the new position
	m.antPos = antNextPos
}
