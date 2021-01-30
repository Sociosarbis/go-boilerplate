package water

import (
	"testing"
)

type Suite struct {
	grid [][]int
	ret  int
}

func TestWater(t *testing.T) {
	suites := []Suite{
		{
			[][]int{{0, 2}, {1, 3}},
			3,
		},
		{
			[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}},
			16,
		},
		{
			[][]int{{3, 2}, {0, 1}},
			3,
		},
	}

	for _, s := range suites {
		ret := swimInWater(s.grid)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
