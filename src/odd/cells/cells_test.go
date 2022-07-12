package cells

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	m       int
	n       int
	indices [][]int
	ret     int
}

func TestCelss(t *testing.T) {
	suites := []suite{
		{
			2,
			3,
			[][]int{{0, 1}, {1, 1}},
			6,
		},
		{
			2,
			2,
			[][]int{{1, 1}, {0, 0}},
			0,
		},
	}

	for _, s := range suites {
		ret := oddCells(s.m, s.n, s.indices)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
