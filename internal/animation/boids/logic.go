package boids

import (
	"math/rand"
)

// initialise the sim
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
	for range numBoids {
		// Start with random positions and velocities
		b := &boid{
			pos: vec2{rand.Float64() * float64(m.Width), rand.Float64() * float64(m.Height)},
			dir: vec2{(rand.Float64()*2 - 1) * maxVelocity, (rand.Float64()*2 - 1) * maxVelocity},
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
			if distance < maxRange {
				repulsion := separation(boid1, boid2, distance)
				boid1.dir = boid1.dir.add(repulsion)
			}
			if distance < neighborDist {
				avgPoss = avgPoss.add(boid2.pos)
				avgDir = avgDir.add(boid2.dir)
				avgLen++
			}
		}

		cohesionForce := cohesion(boid1, avgPoss, avgLen)
		alignmentForce := alignment(avgDir, avgLen)
		boid1.dir = boid1.dir.add(cohesionForce).add(alignmentForce)
		boid1.dir = limitVelocity(boid1.dir)
		boid1.pos = boid1.pos.add(boid1.dir)
		m.edgeCollision(boid1)
	}
}

func (m *Boids) edgeCollision(b *boid) {
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

func limitVelocity(v vec2) vec2 {
	if v.len() > maxVelocity {
		return v.unit().scaled(maxVelocity)
	}
	return v
}

func separation(boid1, boid2 *boid, distance float64) vec2 {
	differenceVector := vec2{(boid1.pos.x - boid2.pos.x), (boid1.pos.y - boid2.pos.y)}
	if distance < minDistance {
		normalVec := normalize(differenceVector)
		repulsion := normalVec.scaled((minDistance - distance) / distance)
		return repulsion
	}
	return vec2{0, 0}
}

func cohesion(boid *boid, avgPos vec2, avgLen int) vec2 {
	if avgLen == 0 {
		return vec2{0, 0}
	}

	avgMass := avgPos.scaled(1.0 / float64(avgLen))

	// Calculate the vector towards the average position
	directionToCenter := avgMass.sub(boid.pos)

	// Normalize and scale the vector (optional: adjust the scaling factor to control strength)
	normalVec := normalize(directionToCenter)

	return normalVec.scaled(cohesionStrength)
}

func alignment(avgDir vec2, avgLen int) vec2 {
	if avgLen == 0 {
		return vec2{0, 0}
	}

	avgDir = avgDir.scaled(1.0 / float64(avgLen))

	normalVec := normalize(avgDir)

	return normalVec.scaled(alignmentStrength)
}
