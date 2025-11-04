package pipes

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	Background(lipgloss.ANSIColor(0))

var bgRender = baseStyle.Render(" ")

var availableColours = []int{
	9,
	10,
	11,
	12,
	13,
	14,
}

var availableSegments = []rune{
	'╗',
	'╔',
	'║',
	'╝',
	'╚',
	'═',
}
var pipeColoursMap = map[segmentKey]string{}

type segmentKey struct {
	color   int
	segment rune
}

var pipeSegmentMap = map[dirKey]rune{
	// Straight Segments
	{vac2{0, 1}, vac2{0, 1}}:   '║', // Up -> Up
	{vac2{0, -1}, vac2{0, -1}}: '║', // Down -> Down
	{vac2{1, 0}, vac2{1, 0}}:   '═', // Right -> Right
	{vac2{-1, 0}, vac2{-1, 0}}: '═', // Left -> Left

	// Corner Segments
	// L-to-R: Vertical segments
	{vac2{0, 1}, vac2{1, 0}}:   '╚', // Up -> Right
	{vac2{0, 1}, vac2{-1, 0}}:  '╝', // Up -> Left
	{vac2{0, -1}, vac2{1, 0}}:  '╔', // Down -> Right
	{vac2{0, -1}, vac2{-1, 0}}: '╗', // Down -> Left

	// R-to-L: Horizontal segments
	{vac2{1, 0}, vac2{0, 1}}:   '╗', // Right -> Up
	{vac2{1, 0}, vac2{0, -1}}:  '╝', // Right -> Down
	{vac2{-1, 0}, vac2{0, 1}}:  '╔', // Left -> Up
	{vac2{-1, 0}, vac2{0, -1}}: '╚', // Left -> Down
}

func (m *Pipes) setupColours() {
	for _, color := range availableColours {
		for _, segment := range availableSegments {
			segmentStyle := lipgloss.NewStyle().
				Foreground(lipgloss.ANSIColor(color)).
				Inherit(baseStyle)
			pipeColoursMap[segmentKey{color, segment}] = segmentStyle.Render(string(segment))
		}
	}
}

func (m *Pipes) View() string {
	// Don't render if grid isn't initialized yet
	if m.Width == 0 || m.Height == 0 || len(m.grid) == 0 {
		return ""
	}

	var view strings.Builder
	// Pre-allocate with a reasonable capacity
	view.Grow(m.Width * m.Height * 12)

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if char, ok := pipeSegmentMap[m.grid[x][y].dir]; ok {
				color := m.grid[x][y].color
				view.WriteString(pipeColoursMap[segmentKey{color, char}])
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
