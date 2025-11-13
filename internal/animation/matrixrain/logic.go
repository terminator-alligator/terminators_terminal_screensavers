package matrixrain

import (
	"math/rand"
)

func (m *MatrixRain) simInit() {
	// Now initialize brightness with the correct dimensions
	m.brightness = make([][]int, m.Width)
	for i := range m.brightness {
		m.brightness[i] = make([]int, m.Height)
	}

	m.grid = make([][]int, m.Width)
	for i := range m.grid {
		m.grid[i] = make([]int, m.Height)
	}
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			randIndx := rand.Intn(len(availableChr))
			m.grid[x][y] = randIndx
		}
	}
}

func (m *MatrixRain) simUpdate() {
	if rand.Float64() <= m.spawnChance {
		randX := rand.Intn(m.Width)
		if m.brightness[randX][0] == 0 {
			m.brightness[randX][0] = m.trailLength
		}
	}
	for y := m.Height - 1; y >= 0; y-- {
		for x := 0; x < m.Width; x++ {
			if m.brightness[x][y] != 0 {
				if y < m.Height-1 {
					m.brightness[x][y+1] = m.brightness[x][y]
				}
				m.brightness[x][y] = m.brightness[x][y] - 1
			}
		}
	}
}
