package testanimation

import (
	"strings"
)

func (m *TestAnimation) View() string {
	// Don't render if grid isn't initialized yet
	var view strings.Builder
	view.WriteString("TEST")

	return view.String()
}
