package count

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	edges   [][]int
	guesses [][]int
	k       int
	ret     int
}

func TestCount(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 1}, {1, 2}, {1, 3}, {4, 2}},
			[][]int{{1, 3}, {0, 1}, {1, 0}, {2, 4}},
			3,
			3,
		},
		{
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			[][]int{{1, 0}, {3, 4}, {2, 1}, {3, 2}},
			1,
			5,
		},
	}

	for _, s := range suites {
		ret := rootCount(s.edges, s.guesses, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
