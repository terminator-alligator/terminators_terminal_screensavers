package animation

import "main.go/config"

type Base struct {
	Config            config.AppConfig
	Width             int
	Height            int
	TimeScale         float64
	FrameCount        int
	AnimationFinished bool
}
