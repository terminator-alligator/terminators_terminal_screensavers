package boids

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"main.go/internal/animation"
)

const (
	numBoids          = 80
	minDistance       = 2.0
	maxRange          = 10
	neighborDist      = 5
	cohesionStrength  = 0.002
	alignmentStrength = 0.0005
	maxVelocity       = 0.1
)

type vec2 struct {
	x float64
	y float64
}

func (v vec2) len() float64 {
	return math.Hypot(v.x, v.y)
}

func (v vec2) scaled(scalar float64) vec2 {
	return vec2{v.x * scalar, v.y * scalar}
}

func (v vec2) unit() vec2 {
	if v.y == 0 && v.x == 0 {
		return vec2{1, 0}
	}
	return v.scaled(1 / v.len())
}

func (v vec2) sub(v2 vec2) vec2 {
	return vec2{
		v.x - v2.x,
		v.y - v2.y,
	}
}

func (v vec2) add(v2 vec2) vec2 {
	return vec2{v.x + v2.x, v.y + v2.y}
}

func normalize(v vec2) vec2 {
	magnitude := v.len()
	return vec2{v.x / magnitude, v.y / magnitude}
}

func euclideanDistance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}

type Boids struct {
	animation.Base
	grid  [][]int
	flock []*boid
}
type boid struct {
	pos vec2
	dir vec2
}

func (m *Boids) Init() tea.Cmd {
	return nil
}

// New implements the animation.IAnimation interface.
func (m *Boids) New() animation.IAnimation {
	return &Boids{}
}

func (m *Boids) Name() string {
	return "Boids"
}
