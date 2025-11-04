package pipes

import "math/rand"

func (m *Pipes) simInit() {
	m.changDirChance = 0.5
	m.pipeSpawnChance = 0.02

	m.grid = make([][]cellState, m.Width)
	for i := range m.grid {
		m.grid[i] = make([]cellState, m.Height)
	}
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			m.grid[x][y] = cellState{}
		}
	}
	m.testAddPipe()
}

func (m *Pipes) testAddPipe() {
	m.pipeList = append(m.pipeList, &pipe{
		4 + rand.Intn(5),
		availableColours[rand.Intn(len(availableColours)-1)],
		vac2{x: rand.Intn(m.Width), y: rand.Intn(m.Width)},
		m.randVec2(),
		vac2{x: 0, y: 0},
	})
}

func (m *Pipes) randVec2() vac2 {
	if rand.Float64() <= 0.5 {
		if rand.Float64() <= 0.5 {
			return vac2{0, 1}
		} else {
			return vac2{0, -1}
		}
	} else {
		if rand.Float64() <= 0.5 {
			return vac2{1, 0}
		} else {
			return vac2{-1, 0}
		}
	}
}

func (m *Pipes) simUpdate() {
	if rand.Float64() <= m.pipeSpawnChance {
		m.testAddPipe()
	}
	var nextPipeList []*pipe
	for _, p := range m.pipeList {
		p.prevDir = p.dir
		p.pos = vac2{p.pos.x + p.dir.x, p.pos.y + p.dir.y}
		p.moveLength--
		if p.moveLength <= 0 {
			if rand.Float64() <= m.changDirChance {
				if rand.Float64() <= 0.5 {
					m.turnRight(p)
				} else {
					m.turnLeft(p)
				}
			}
			p.moveLength = 4 + rand.Intn(5)
		}
		if !m.bounceCheck(p) {
			m.grid[p.pos.x][p.pos.y] = cellState{dirKey{p.prevDir, p.dir}, p.color}
			nextPipeList = append(nextPipeList, p)
		}
	}
	m.pipeList = nextPipeList
}

func (m *Pipes) turnRight(p *pipe) {
	p.dir = vac2{-p.dir.y, p.dir.x}
}

func (m *Pipes) turnLeft(p *pipe) {
	p.dir = vac2{p.dir.y, -p.dir.x}
}

func (m *Pipes) bounceCheck(p *pipe) bool {
	if p.pos.x < 0 {
		return true
	} else if p.pos.x >= m.Width {
		return true
	}
	if p.pos.y < 0 {
		return true
	} else if p.pos.y >= m.Height {
		return true
	}
	return false
}
