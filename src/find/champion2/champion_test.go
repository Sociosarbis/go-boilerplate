package champion

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	edges [][]int
	ret   int
}

func TestSuite(t *testing.T) {
	suites := []suite{
		{
			3,
			[][]int{{0, 1}, {1, 2}},
			0,
		},
		{
			4,
			[][]int{{0, 2}, {1, 3}, {1, 2}},
			-1,
		},
	}

	for _, s := range suites {
		ret := findChampion(s.n, s.edges)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
