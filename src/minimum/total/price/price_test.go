package price

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n     int
	edges [][]int
	price []int
	trips [][]int
	ret   int
}

func TestPrice(t *testing.T) {
	suites := []suite{
		{
			4,
			[][]int{{0, 1}, {1, 2}, {1, 3}},
			[]int{2, 2, 10, 6},
			[][]int{{0, 3}, {2, 1}, {2, 3}},
			23,
		},
		{
			2,
			[][]int{{0, 1}},
			[]int{2, 2},
			[][]int{{0, 0}},
			1,
		},
	}

	for _, s := range suites {
		ret := minimumTotalPrice(s.n, s.edges, s.price, s.trips)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
