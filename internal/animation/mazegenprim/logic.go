package mazegenprim

import "math/rand"

// initialise the sim
func (m *MazeGenerationPrims) simInit() {
	m.wallsList = [][2]int{}
	m.startPos = [2]int{}
	m.endPos = [2]int{-1, -1}
	// endPos is set to an impossible to avoid rendering it
	m.grid = make([][]gridState, m.Width)
	for i := range m.grid {
		m.grid[i] = make([]gridState, m.Height)
	}
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			m.grid[x][y] = wall
		}
	}
	m.pickStartPos()
	m.addWalls(m.startPos[0], m.startPos[1])
}

func (m *MazeGenerationPrims) pickStartPos() {
	var x, y int
	x = rand.Intn((m.Width - 1))
	y = rand.Intn((m.Height - 1))
	if x%2 == 0 {
		x++
	}
	if y%2 == 0 {
		y++
	}
	m.startPos = [2]int{x, y}
	m.grid[x][y] = open
}

func (m *MazeGenerationPrims) addWalls(x, y int) {
	// Add neighboring walls of cell (x, y) to the list
	for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		wallX, wallY := x+d[0], y+d[1]
		if wallX > 0 && wallX < m.Width-1 && wallY > 0 && wallY < m.Height-1 {
			m.wallsList = append(m.wallsList, [2]int{wallX, wallY})
		}
	}
}

func (m *MazeGenerationPrims) simUpdate() {
	if len(m.wallsList) == 0 {
		m.finishedGenerating = true
		return
	}

	// Pick a random wall from the list
	randWallIndex := rand.Intn(len(m.wallsList))
	wallCoords := m.wallsList[randWallIndex]

	// Remove the wall from the list by swapping with the last element
	x, y := wallCoords[0], wallCoords[1]

	// Find out how many adjacent cells are open passages.
	// A valid wall to break connects an open passage to a still-walled cell.
	openNeighbors := 0
	var oppositeCell [2]int
	for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < m.Width && ny >= 0 && ny < m.Height && m.grid[nx][ny] == open {
			openNeighbors++
			// This is the cell on the other side of the wall, relative to the open passage.
			// It's the one we will open up.
			oppositeCell = [2]int{x - d[0], y - d[1]}
		}
	}

	// If the wall separates an open cell from a walled-off cell, open the wall.
	// We also need to check if the opposite cell is within bounds and is currently a wall.
	if openNeighbors == 1 &&
		oppositeCell[0] >= 0 && oppositeCell[0] < m.Width && oppositeCell[1] >= 0 && oppositeCell[1] < m.Height && m.grid[oppositeCell[0]][oppositeCell[1]] == wall {
		m.grid[x][y] = open
		m.grid[oppositeCell[0]][oppositeCell[1]] = open
		m.addWalls(oppositeCell[0], oppositeCell[1])
	}

	m.wallsList[randWallIndex] = m.wallsList[len(m.wallsList)-1]
	m.wallsList = m.wallsList[:len(m.wallsList)-1]

	if len(m.wallsList) == 0 {
		m.endPos = [2]int{x, y}
	}
}
