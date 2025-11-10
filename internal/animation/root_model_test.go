package animation_test

import (
	"testing"

	"main.go/config"
	"main.go/internal/animation"
)

func TestRootModel_NextAnim(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		AppConfig    config.AppConfig
		initialAnnim string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := animation.NewRootModel(tt.AppConfig, tt.initialAnnim)
			m.NextAnim()
		})
	}
}
