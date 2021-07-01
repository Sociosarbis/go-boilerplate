package ways

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n        int
	relation [][]int
	k        int
	ret      int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			5,
			[][]int{{0, 2}, {2, 1}, {3, 4}, {2, 3}, {1, 4}, {2, 0}, {0, 4}},
			3,
			3,
		},
		{
			3,
			[][]int{{0, 2}, {2, 1}},
			2,
			0,
		},
	}

	for _, s := range suites {
		ret := numWays(s.n, s.relation, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
