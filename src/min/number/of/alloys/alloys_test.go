package alloys

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n           int
	k           int
	budget      int
	composition [][]int
	stock       []int
	cost        []int
	ret         int
}

func TestAlloys(t *testing.T) {
	suites := []suite{
		{
			3,
			2,
			15,
			[][]int{{1, 1, 1}, {1, 1, 10}},
			[]int{0, 0, 0},
			[]int{1, 2, 3},
			2,
		},
		{
			3,
			2,
			15,
			[][]int{{1, 1, 1}, {1, 1, 10}},
			[]int{0, 0, 100},
			[]int{1, 2, 3},
			5,
		},
		{
			2,
			3,
			10,
			[][]int{{2, 1}, {1, 2}, {1, 1}},
			[]int{1, 1},
			[]int{5, 5},
			2,
		},
	}

	for _, s := range suites {
		ret := maxNumberOfAlloys(s.n, s.k, s.budget, s.composition, s.stock, s.cost)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
