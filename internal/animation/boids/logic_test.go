package boids

import (
	"reflect"
	"testing"

	"main.go/internal/animation"
)

func TestBoids_edgeCollision(t *testing.T) {
	tests := []struct {
		name        string // description of this test case
		width       int
		height      int
		initialPos  vec2
		initialDir  vec2
		expectedDir vec2
		expectedPos vec2
	}{
		{
			name:        "top collision",
			width:       20,
			height:      20,
			initialPos:  vec2{10, -5},
			initialDir:  vec2{0, -10},
			expectedPos: vec2{10, 0},
			expectedDir: vec2{0, 10},
		},
		{
			name:        "bottom collision",
			width:       20,
			height:      20,
			initialPos:  vec2{10, 25},
			initialDir:  vec2{0, 10},
			expectedPos: vec2{10, 19},
			expectedDir: vec2{0, -10},
		},
		{
			name:        "left collision",
			width:       20,
			height:      20,
			initialPos:  vec2{-5, 10},
			initialDir:  vec2{-10, 0},
			expectedPos: vec2{0, 10},
			expectedDir: vec2{10, 0},
		},
		{
			name:        "right collision",
			width:       20,
			height:      20,
			initialPos:  vec2{25, 10},
			initialDir:  vec2{10, 0},
			expectedPos: vec2{19, 10},
			expectedDir: vec2{-10, 0},
		},
		{
			name:        "no collision",
			width:       20,
			height:      20,
			initialPos:  vec2{10, 10},
			initialDir:  vec2{0, 0},
			expectedPos: vec2{10, 10},
			expectedDir: vec2{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Boids{Base: animation.Base{Width: tt.width, Height: tt.height}}
			b := &boid{
				pos: tt.initialPos,
				dir: tt.initialDir,
			}
			m.edgeCollision(b)

			if !reflect.DeepEqual(b.pos, tt.expectedPos) {
				t.Errorf("edgeCollision() pos = %v, want %v", b.pos, tt.expectedPos)
			}
			if !reflect.DeepEqual(b.dir, tt.expectedDir) {
				t.Errorf("edgeCollision() dir = %v, want %v", b.dir, tt.expectedDir)
			}
		})
	}
}

func Test_limitVelocity(t *testing.T) {
	const maxVelocity = 2.0
	tests := []struct {
		name            string
		initialVelocity vec2
	}{
		{
			name:            "velocity above max",
			initialVelocity: vec2{maxVelocity + 10, maxVelocity + 10},
		},
		{
			name:            "velocity below max",
			initialVelocity: vec2{maxVelocity - 1, maxVelocity - 1},
		},
		{
			name:            "zero velocity",
			initialVelocity: vec2{0, 0},
		},
		{
			name:            "negative velocity",
			initialVelocity: vec2{-maxVelocity - 10, -maxVelocity - 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Boids{maxVelocity: maxVelocity}
			got := m.limitVelocity(tt.initialVelocity)
			if got.len() > maxVelocity {
				t.Errorf("limitVelocity() = %v, length %v, expected max length %v", got, got.len(), maxVelocity)
			} else if got.len() < 0 {
				t.Errorf("limitVelocity() = %v, length %v, expected non-negative length", got, got.len())
			}
		})
	}
}
