package orders

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	orders [][]int
	ret    int
}

func TestOrders(t *testing.T) {
	suites := []suite{
		{
			[][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}},
			6,
		},
		{
			[][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}},
			999999984,
		},
	}

	for _, s := range suites {
		ret := getNumberOfBacklogOrders(s.orders)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
