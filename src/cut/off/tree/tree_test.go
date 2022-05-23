package tree

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	forest [][]int
	ret    int
}

func TestTree(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}},
			6,
		},
		{
			[][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}},
			-1,
		},
		{
			[][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}},
			6,
		},
	}

	for _, s := range suites {
		ret := cutOffTree(s.forest)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
