package template

import "sort"

type MineSorter struct {
	scores []int
}

func (m MineSorter) Len() int {
	return len(m.scores)
}

func (m MineSorter) Less(i, j int) bool {
	return m.scores[i] < m.scores[j]
}

func (m MineSorter) Swap(i, j int) {
	m.scores[i], m.scores[j] = m.scores[j], m.scores[i]
}

func Example() {
	sorter := MineSorter{scores: []int{4, 2, 2, 5, 3, 9, 0, 1}}
	sort.Sort(sorter)
}
