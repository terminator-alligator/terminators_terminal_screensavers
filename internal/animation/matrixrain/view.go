package matrixrain

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

var bgRender = baseStyle.Render(" ")

type chrKey struct {
	brightness int
	chrIndex   int
}

var renderChrMap = map[chrKey]string{}

func (m *MatrixRain) setUpChrMap() {
	brightnessMultiplier := 256 / m.trailLength
	for index, chrRune := range availableChr {
		// start from 1 because a brightness of zero is not used
		for brightnessLevel := 1; brightnessLevel <= m.trailLength; brightnessLevel++ {
			var r, g, b uint8
			// make the last character brighter
			if brightnessLevel == m.trailLength {
				r = 200
				b = 200
			}
			g = uint8(brightnessMultiplier * brightnessLevel)
			g = max(g, 50)
			key := chrKey{brightness: brightnessLevel, chrIndex: index}
			chrStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color(rgbToHex(r, g, b))).
				Inherit(baseStyle)
			renderChrMap[key] = chrStyle.Render(string(chrRune))
		}
	}
}

func rgbToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (m *MatrixRain) View() string {
	if m.Width == 0 || m.Height == 0 || len(m.grid) == 0 {
		return ""
	}

	var view strings.Builder
	// Pre-allocate with a reasonable capacity
	view.Grow(m.Width * m.Height * 12)

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.brightness[x][y] != 0 {
				key := chrKey{brightness: m.brightness[x][y], chrIndex: m.grid[x][y]}
				view.WriteString(renderChrMap[key])
			} else {
				view.WriteString(bgRender)
			}
		}
		if y < m.Height-1 {
			view.WriteRune('\n')
		}
	}
	return view.String()
}
