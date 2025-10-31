package langtonsant

import (
	"reflect"
	"testing"

	"main.go/internal/animation"
)

func TestLangtonsAnt_simInit(t *testing.T) {
	tests := []struct {
		name        string
		width       int
		height      int
		expectedPos vec2
		expectedDir vec2
	}{
		{
			name:        "10x10 grid",
			width:       10,
			height:      10,
			expectedPos: vec2{x: 5, y: 5},
			expectedDir: vec2{x: 0, y: -1},
		},
		{
			name:        "20x20 grid",
			width:       20,
			height:      20,
			expectedPos: vec2{x: 10, y: 10},
			expectedDir: vec2{x: 0, y: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := LangtonsAnt{Base: animation.Base{Width: tt.width, Height: tt.height}}
			m.simInit()

			if !reflect.DeepEqual(m.antPos, tt.expectedPos) {
				t.Errorf("simInit() antPos = %v, want %v", m.antPos, tt.expectedPos)
			}
			if !reflect.DeepEqual(m.antDir, tt.expectedDir) {
				t.Errorf("simInit() antDir = %v, want %v", m.antDir, tt.expectedDir)
			}
		})
	}
}

func TestLangtonsAnt_simUpdate(t *testing.T) {
	tests := []struct {
		name        string
		width       int
		height      int
		expectedPos vec2
		expectedDir vec2
	}{
		{
			name:        "ant on dead cell turns right and moves forward",
			width:       20,
			height:      20,
			expectedPos: vec2{x: 10, y: 9},
			expectedDir: vec2{x: 1, y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := LangtonsAnt{Base: animation.Base{Width: tt.width, Height: tt.height}}
			m.simInit()
			m.simUpdate()

			if !reflect.DeepEqual(m.antPos, tt.expectedPos) {
				t.Errorf("simUpdate() antPos = %v, want %v", m.antPos, tt.expectedPos)
			}
			if !reflect.DeepEqual(m.antDir, tt.expectedDir) {
				t.Errorf("simUpdate() antDir = %v, want %v", m.antDir, tt.expectedDir)
			}
		})
	}
}
