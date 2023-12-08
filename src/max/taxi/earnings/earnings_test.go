package earnings

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	rides [][]int
	ret   int64
}

func TestEarnings(t *testing.T) {
	suites := []suite{
		{
			5,
			[][]int{{2, 5, 4}, {1, 5, 1}},
			7,
		},
		{
			20,
			[][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}},
			20,
		},
		{
			11,
			[][]int{{1, 10, 5000}, {2, 6, 14}, {3, 6, 155}},
			5009,
		},
	}

	for _, s := range suites {
		ret := maxTaxiEarnings(s.n, s.rides)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
