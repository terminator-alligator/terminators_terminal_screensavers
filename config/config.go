package config

import (
	"os"
	"path"

	toml "github.com/pelletier/go-toml/v2"
)

const (
	AppName        = "ttss"
	ConfigFileName = "config.toml"
)

// AppConfig holds the entire configuration structure.
type AppConfig struct {
	Global         GlobalConfig         `toml:"global"`
	Boids          BoidsConfig          `toml:"boids"`
	BobbleSort     BobbleSortConfig     `toml:"bobble_sort"`
	LangtonsAnt    LangtonsAntConfig    `toml:"langtons_ant"`
	MazeGeneration MazeGenerationConfig `toml:"maze_generation"`
	Pipes          PipesConfig          `toml:"pipes"`
}

type GlobalConfig struct {
	FrameRate          float64  `toml:"frame_rate"`
	DebugMode          bool     `toml:"debug_mode"`
	SelectedAnimations []string `toml:"selected_animations"`
	Shuffle            bool     `toml:"shuffle_mode"`
}

type BobbleSortConfig struct {
	TimeScale float64 `toml:"time_scale"`
}

type MazeGenerationConfig struct {
	TimeScale float64 `toml:"time_scale"`
}
type LangtonsAntConfig struct {
	TimeScale float64 `toml:"time_scale"`
}

type PipesConfig struct {
	TimeScale       float64 `toml:"time_scale"`
	ChangDirChance  float64 `toml:"chang_direction_chance"`
	PipeSpawnChance float64 `toml:"pipe_spawn_chance"`
}

type BoidsConfig struct {
	TimeScale        float64 `toml:"time_scale"`
	NumBoids         int     `toml:"num_boids"`
	MaxVelocity      float64 `toml:"max_velocity"`
	NeighborDist     float64 `toml:"neighbor_dist"`
	MaxRange         float64 `toml:"max_Range"`
	CohesionWeight   float64 `toml:"cohesion_weight"`
	AlignmentWeight  float64 `toml:"alignment_weight"`
	SeparationWeight float64 `toml:"separation_weight"`
}

func NewDefaultConfig() AppConfig {
	return AppConfig{
		Global: GlobalConfig{
			FrameRate:          60.0,
			DebugMode:          false,
			SelectedAnimations: []string{"LangtonsAnt", "BubbleSort", "MazeGenerationPrims", "Boids", "Pipes"},
			Shuffle:            true,
		},
		Boids: BoidsConfig{
			TimeScale:        1.0,
			NumBoids:         40,
			MaxRange:         10.0,
			NeighborDist:     10.0,
			MaxVelocity:      0.1,
			CohesionWeight:   0.002,
			AlignmentWeight:  0.0005,
			SeparationWeight: 2.0,
		},
		BobbleSort: BobbleSortConfig{
			TimeScale: 1.0,
		},
		LangtonsAnt: LangtonsAntConfig{
			TimeScale: 1.0,
		},
		MazeGeneration: MazeGenerationConfig{
			TimeScale: 1.0,
		},
		Pipes: PipesConfig{
			TimeScale:       1.0,
			ChangDirChance:  0.5,
			PipeSpawnChance: 0.02,
		},
	}
}

func Load() (AppConfig, error) {
	config := NewDefaultConfig()
	// load the defaults first

	// TODO: this still needs a lot of work

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return AppConfig{}, err
	}
	configPath := path.Join(homeDir, ".config", AppName, ConfigFileName)

	data, err := os.ReadFile(configPath)
	if err != nil {
		if err == os.ErrNotExist {
			// If the config file doesn't exist, return the default config
			// TODO: instead of this we should create the file
			return config, nil
		}
		return AppConfig{}, err
	}

	// override the defaults with the users
	if err := toml.Unmarshal(data, &config); err != nil {
		return AppConfig{}, err
	}
	return config, nil
}
