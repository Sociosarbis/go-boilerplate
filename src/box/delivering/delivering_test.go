package delivering

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	boxes      [][]int
	portsCount int
	maxBoxes   int
	maxWeight  int
	ret        int
}

func TestDelivering(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 1}, {2, 1}, {1, 1}},
			2,
			3,
			3,
			4,
		},
		{
			[][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}},
			3,
			3,
			6,
			6,
		},
		{
			[][]int{{1, 4}, {1, 2}, {2, 1}, {2, 1}, {3, 2}, {3, 4}},
			3,
			6,
			7,
			6,
		},
		{
			[][]int{{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2}},
			5,
			5,
			7,
			14,
		},
	}

	for _, s := range suites {
		ret := boxDelivering(s.boxes, s.portsCount, s.maxBoxes, s.maxWeight)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
