package algorithm

import "testing"

type testpair struct {
	source []int
	sorted []int
}

var tests = []testpair{
	{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{2, 1, 3}, []int{1, 2, 3}},
	{[]int{5, 4, 2}, []int{2, 4, 5}},
}

type funcSort func(data []int)

func sorted(t *testing.T, f funcSort) {
	for _, pair := range tests {
		tmp := make([]int, len(pair.source))
		copy(tmp, pair.source)

		f(tmp)

		if !Equal(tmp, pair.sorted) {
			t.Error(
				"For", pair.source,
				"expected", pair.sorted,
				"got", tmp)
		}
	}
}
