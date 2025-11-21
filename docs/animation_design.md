# Animation Design Guide

This guide provides the foundational knowledge needed to create new animations for the Terminators Terminal Screensavers project. It outlines the core philosophy, required structure, and interface that all animations must follow.

While this document aims to be comprehensive, it is a living document and may evolve as the project grows.

## Code Style and Quality

This project does not enforce a strict style guide beyond standard Go conventions (e.g., `gofmt`). However, all contributions should prioritize readability and maintainability.

Since animations are self-contained modules, there is room for creative implementation. The primary goal is to produce clean, understandable code that others can learn from and build upon.

## Animation Design Philosophy

*   Animations should aim to be visually appealing or represent a cool concept.
*   Animations should be procedurally generated; at this moment there is no desire to add regular ASCII art.
*   Animations should try to not be resource intensive.
*   Animations should try and contain comments where necessary.


## Animation Interface

All animations must implement the `IAnimation` interface:

```go
type IAnimation interface {
    Init() tea.Cmd
    Name() string
    GetTimeScale() float64
    Update(teaMsg tea.Msg) (IAnimation, tea.Cmd)
    New(config.AppConfig) IAnimation
    View() string
}
```

### Method Descriptions

*   `Init()`: Initializes the animation.
*   `Name()`: Returns the name of the animation.
*   `GetTimeScale()`: Returns the time scale factor for the animation speed.
*   `Update(tea.Msg)`: Handles updates.
*   `New(config.AppConfig)`: Creates a new instance of the animation with the given configuration.
*   `View()`: Returns the current state of the animation as a string.

## Core Concepts

Animations are structured into four main files:

*   **logic.go**: This file is where all logic relevant to the animation should be.
*   **model.go**: This should contain information about the model, any relevant types or structures, and any documentation relevant to the animation.
*   **update.go**: This is where updates are received from the root animation. It can contain logic for switching animations when the current one has finished.
*   **view.go**: This is where all logic relevant to rendering the animation should be located. If a structure or type is not used anywhere else and is relevant only to rendering, it should be located here.

### `logic.go`

This file should contain the core simulation logic, separated from the Bubble Tea model.

*   **`simInit()`**: Contains logic for initializing and resetting the simulation state.
*   **`simUpdate()`**: Contains logic for updating the simulation. Each call to this function represents a single frame, and each frame should aim to be different from the last.

### `model.go`

This file defines the animation's main `Model` struct and implements several methods of the `IAnimation` interface.

*   **The `Model` struct**: This struct contains all the information and state for the animation.
*   **`Init()`**: Calls any functions required for the initial setup of the animation model.
*   **`New()`**: Sets up and returns a new instance of the animation with the given configuration applied.
*   **`GetTimeScale()`**: Returns the time scale value, typically from a base animation struct.
*   **`Name()`**: Returns the animation's name. This is used to list and select the animation to run.

### `view.go`

This file should contain a single primary function.

*   **`View()`**: Contains the logic for rendering the current state of the animation model into a string for display.

### `update.go`

This file should also contain a single primary function.

*   **`Update()`**: Contains the logic for handling incoming messages (`tea.Msg`). This includes handling window resize events and determining what to do when the animation has finished.
