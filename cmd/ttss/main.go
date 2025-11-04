package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	tea "github.com/charmbracelet/bubbletea"
	"main.go/config"
	"main.go/internal/animation"
	"main.go/internal/animation/boids"
	"main.go/internal/animation/bubblesort"
	"main.go/internal/animation/langtonsant"
	"main.go/internal/animation/mazegenprim"
	"main.go/internal/animation/pipes"
)

// using flags for now
var (
	listAnimations = flag.Bool("list", false, "List available animations")
	runAnimation   = flag.String("run", "", "Run a specific animation by name (e.g., 'Langton's Ant')")
)

func main() {
	flag.Parse()
	config, err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	AvailableAnimations := []animation.IAnimation{
		(&langtonsant.LangtonsAnt{}).New(config),
		(&bubblesort.BubbleSort{}).New(config),
		(&mazegenprim.MazeGenerationPrims{}).New(config),
		(&boids.Boids{}).New(config),
		(&pipes.Pipes{}).New(config),
	}

	// set up the logger
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer func() {
		_ = f.Close()
	}()
	log.Println("Application started")

	// CPU Profiling
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer func() {
		_ = cpuFile.Close()
	}()
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	var selectedAnimation animation.IAnimation

	if *listAnimations {
		for _, anim := range AvailableAnimations {
			fmt.Println(anim.Name())
		}
		os.Exit(0)
	}

	if *runAnimation != "" {
		for _, anim := range AvailableAnimations {
			if anim.Name() == *runAnimation {
				selectedAnimation = anim
				break
			}
		}
		// If no animation was found with the given name, exit with an error
		if selectedAnimation == nil {
			fmt.Printf("Error: Animation '%s' not found.\n", *runAnimation)
			os.Exit(1)
		}
	} else {
		// Default to the third animation (Maze Generation) if no specific animation is requested
		selectedAnimation = AvailableAnimations[4]
	}

	m := animation.NewRootModel(config, selectedAnimation)

	if _, err := tea.NewProgram(&m).Run(); err != nil {
		fmt.Println("Error running program", err)
		os.Exit(1)
	}
}
