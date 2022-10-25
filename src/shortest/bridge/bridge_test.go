package bridge

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	ret  int
}

func TestBridge(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 1}, {1, 0}},
			1,
		},
		{
			[][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}},
			2,
		},
		{
			[][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}},
			1,
		},
	}

	for _, s := range suites {
		ret := shortestBridge(s.grid)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
