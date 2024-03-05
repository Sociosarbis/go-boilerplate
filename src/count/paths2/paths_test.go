package paths2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	roads [][]int
	ret   int
}

func TestPaths(t *testing.T) {
	suites := []suite{
		{
			7,
			[][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}},
			4,
		},
		{
			2,
			[][]int{{1, 0, 10}},
			1,
		},
	}

	for _, s := range suites {
		ret := countPaths(s.n, s.roads)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
