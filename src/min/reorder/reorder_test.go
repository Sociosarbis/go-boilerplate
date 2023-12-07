package reorder

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n           int
	connections [][]int
	ret         int
}

func TestReorder(t *testing.T) {
	suites := []suite{
		{
			6,
			[][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}},
			3,
		},
		{
			5,
			[][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}},
			2,
		},
		{
			3,
			[][]int{{1, 0}, {2, 0}},
			0,
		},
	}

	for _, s := range suites {
		ret := minReorder(s.n, s.connections)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
