package cost

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	roads [][]int
	seats int
	ret   int64
}

func TestCost(t *testing.T) {
	suites := []suite{
		{
			[][]int{{0, 1}, {0, 2}, {0, 3}},
			5,
			3,
		},
		{
			[][]int{{3, 1}, {3, 2}, {1, 0}, {0, 4}, {0, 5}, {4, 6}},
			2,
			7,
		},
		{
			[][]int{},
			1,
			0,
		},
		{
			[][]int{{1, 0}, {1, 2}, {0, 3}, {4, 3}, {5, 2}, {4, 6}, {1, 7}, {8, 6}, {9, 6}, {1, 10}, {6, 11}, {1, 12}, {13, 9}, {4, 14}, {3, 15}, {2, 16}, {5, 17}, {3, 18}, {6, 19}, {6, 20}, {21, 16}, {18, 22}, {0, 23}, {24, 1}, {25, 12}, {26, 24}, {9, 27}, {28, 23}, {29, 25}, {6, 30}, {31, 21}, {32, 21}, {33, 23}, {19, 34}, {5, 35}, {36, 7}, {25, 37}, {0, 38}, {1, 39}, {6, 40}, {41, 3}},
			5,
			52,
		},
	}

	for _, s := range suites {
		ret := minimumFuelCost(s.roads, s.seats)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
