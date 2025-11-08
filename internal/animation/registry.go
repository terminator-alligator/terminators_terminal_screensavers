package animation

import (
	"main.go/config"
	"main.go/internal/animation/base"
	"main.go/internal/animation/boids"
	"main.go/internal/animation/bubblesort"
	"main.go/internal/animation/langtonsant"
	"main.go/internal/animation/mazegenprim"
	"main.go/internal/animation/pipes"
)

func GetAvailableAnimations(config config.AppConfig) []base.IAnimation {
	return []base.IAnimation{
		(&langtonsant.LangtonsAnt{}).New(config),
		(&bubblesort.BubbleSort{}).New(config),
		(&mazegenprim.MazeGenerationPrims{}).New(config),
		(&boids.Boids{}).New(config),
		(&pipes.Pipes{}).New(config),
	}
}
