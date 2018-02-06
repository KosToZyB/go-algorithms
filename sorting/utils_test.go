package algorithm

import "testing"

type equalPair struct {
	source []int
	sorted []int
}

var pasitiveScenario = []equalPair{
	{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
	{[]int{2, 1, 3}, []int{2, 1, 3}},
	{[]int{5}, []int{5}},
	{nil, nil},
}

var negativeScenario = []equalPair{
	{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2}},
	{[]int{5, 4, 3, 1}, []int{5, 4, 3, 2}},
	{[]int{2, 1, 3}, []int{}},
	{[]int{5}, []int{5, 1}},
	{[]int{5}, nil},
}

func TestEqual(t *testing.T) {
	for _, pair := range pasitiveScenario {
		if !Equal(pair.sorted, pair.source) {
			t.Error(
				"For", pair.source,
				"and", pair.sorted,
				"expected true",
				"got false")
		}
	}

	for _, pair := range negativeScenario {
		if Equal(pair.sorted, pair.source) {
			t.Error(
				"For", pair.source,
				"and", pair.sorted,
				"expected false",
				"got true")
		}
	}
}
