package bubblesort

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	barStyle = lipgloss.NewStyle().Background(lipgloss.Color("5"))
	// emptyStyle = lipgloss.NewStyle().Background(lipgloss.Color("236"))
	emptyStyle = lipgloss.NewStyle().Background(lipgloss.ANSIColor(0))
)

func (m *BubbleSort) View() string {
	if m.Width == 0 || m.Height == 0 || len(m.items) == 0 {
		return ""
	}

	numItems := len(m.items)
	barWidth := m.Width / numItems
	// Ensure barWidth is at least 1 if there are items to display
	if barWidth == 0 && numItems > 0 {
		barWidth = 1
	}
	// Calculate remaining space to distribute
	remainingSpace := m.Width - (barWidth * numItems)

	var view strings.Builder
	view.Grow(m.Width * m.Height) // A more accurate pre-allocation

	maxValue := 0
	for _, item := range m.items {
		if item > maxValue {
			maxValue = item
		}
	}
	// Avoid division by zero if all items are 0
	if maxValue == 0 {
		maxValue = 1
	}

	for y := 0; y < m.Height; y++ {
		for i, item := range m.items {
			barHeight := (item * m.Height) / maxValue
			style := emptyStyle
			if m.Height-y <= barHeight {
				if i == m.curentIndx {
					style = barStyle.Background(lipgloss.ANSIColor(1))
				} else {
					style = barStyle.Background(lipgloss.ANSIColor(item))
				}
			}
			view.WriteString(strings.Repeat(style.Render(" "), barWidth))
			// Distribute remaining space one character at a time to the first few bars
			if i < remainingSpace {
				view.WriteString(style.Render(" "))
			}
		}
		if y < m.Height-1 {
			view.WriteRune('\n')
		}
	}

	return view.String()
}
