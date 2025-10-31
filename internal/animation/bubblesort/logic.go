package bubblesort

import "math/rand"

func (m *BubbleSort) simInit() {
	m.items = []int{}
	for i := range 25 {
		m.items = append(m.items, i)
	}
	rand.Shuffle(len(m.items), func(i int, j int) {
		m.items[i], m.items[j] = m.items[j], m.items[i]
	})
}

func (m *BubbleSort) simUpdate() {
	if len(m.items) < 2 {
		return
	}
	if m.curentIndx >= len(m.items)-1 {
		m.curentIndx = -1 // Reset for the next pass
	} else if m.items[m.curentIndx] > m.items[m.curentIndx+1] {
		m.items[m.curentIndx], m.items[m.curentIndx+1] = m.items[m.curentIndx+1], m.items[m.curentIndx]
	}
	m.curentIndx++
}
