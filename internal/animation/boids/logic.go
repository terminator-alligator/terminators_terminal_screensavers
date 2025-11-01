package boids

import (
	"math/rand"
)

func (m *Boids) simInit() {
	m.flock = []*boid{}
	m.initBoids()

	m.grid = make([][]int, m.Width)
	for i := range m.grid {
		m.grid[i] = make([]int, m.Height)
	}
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			m.grid[x][y] = 0
		}
	}
}

func (m *Boids) initBoids() {
	for range m.numBoids {
		// Start with random positions and velocities
		b := &boid{
			pos: vec2{rand.Float64() * float64(m.Width), rand.Float64() * float64(m.Height)},
			dir: vec2{(rand.Float64()*2 - 1) * m.maxVelocity, (rand.Float64()*2 - 1) * m.maxVelocity},
		}
		m.flock = append(m.flock, b)
	}
}

func (m *Boids) simUpdate() {
	// Reset grid
	for i := range m.grid {
		for j := range m.grid[i] {
			m.grid[i][j] = 0
		}
	}

	for _, boid1 := range m.flock {
		avgPoss := vec2{0, 0}
		avgDir := vec2{0, 0}
		avgLen := 0
		for _, boid2 := range m.flock {
			if boid1 == boid2 {
				continue
			}
			distance := euclideanDistance(boid1.pos.x, boid1.pos.y, boid2.pos.x, boid2.pos.y)
			if distance < m.maxRange {
				repulsion := m.separation(boid1, boid2, distance)
				boid1.dir = boid1.dir.add(repulsion)
			}
			if distance < m.neighborDist {
				avgPoss = avgPoss.add(boid2.pos)
				avgDir = avgDir.add(boid2.dir)
				avgLen++
			}
		}

		cohesionForce := m.cohesion(boid1, avgPoss, avgLen)
		alignmentForce := m.alignment(avgDir, avgLen)
		boid1.dir = boid1.dir.add(cohesionForce).add(alignmentForce)
		boid1.dir = m.limitVelocity(boid1.dir)
		boid1.pos = boid1.pos.add(boid1.dir)
		m.edgeCollision(boid1)
	}
}

func (m *Boids) edgeCollision(b *boid) {
	// flip the direction
	if b.pos.x < 0 {
		b.pos.x = 0
		b.dir.x = -b.dir.x
	} else if b.pos.x >= float64(m.Width) {
		b.pos.x = float64(m.Width - 1)
		b.dir.x = -b.dir.x
	}
	if b.pos.y < 0 {
		b.pos.y = 0
		b.dir.y = -b.dir.y
	} else if b.pos.y >= float64(m.Height) {
		b.pos.y = float64(m.Height - 1)
		b.dir.y = -b.dir.y
	}
}

func (m *Boids) limitVelocity(v vec2) vec2 {
	if v.len() > m.maxVelocity {
		return v.unit().scaled(m.maxVelocity)
	}
	return v
}

func (m *Boids) separation(boid1, boid2 *boid, distance float64) vec2 {
	differenceVector := vec2{(boid1.pos.x - boid2.pos.x), (boid1.pos.y - boid2.pos.y)}
	if distance < m.minDistance {
		normalVec := normalize(differenceVector)
		repulsion := normalVec.scaled((m.minDistance - distance) / distance)
		return repulsion
	}
	return vec2{0, 0}
}

func (m *Boids) cohesion(boid *boid, avgPos vec2, avgLen int) vec2 {
	if avgLen == 0 {
		return vec2{0, 0}
	}

	avgMass := avgPos.scaled(1.0 / float64(avgLen))

	// Calculate the vector towards the average position
	directionToCenter := avgMass.sub(boid.pos)

	// Normalize and scale the vector
	normalVec := normalize(directionToCenter)

	return normalVec.scaled(m.cohesionWeight)
}

func (m *Boids) alignment(avgDir vec2, avgLen int) vec2 {
	if avgLen == 0 {
		return vec2{0, 0}
	}

	avgDir = avgDir.scaled(1.0 / float64(avgLen))

	normalVec := normalize(avgDir)

	return normalVec.scaled(m.alignmentWeight)
}
