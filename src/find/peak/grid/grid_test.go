package grid

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	mat [][]int
	ret []int
}

func TestGrid(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 4}, {3, 2}},
			[]int{0, 1},
		},
		{
			[][]int{{2, 1, 99, 98, 22, 21, 20}, {3, 2, 1, 2, 3, 4, 19}, {4, 3, 2, 1, 2, 3, 18}, {5, 4, 3, 2, 1, 2, 17}, {6, 5, 4, 3, 2, 1, 16}, {7, 6, 5, 4, 3, 2, 15}, {8, 9, 10, 11, 12, 13, 14}},
			[]int{0, 2},
		},
		{
			[][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}},
			[]int{2, 2},
		},
		{
			[][]int{{10, 20, 40, 50, 60, 70}, {1, 4, 2, 3, 500, 80}},
			[]int{1, 4},
		},
		{
			[][]int{{7, 2, 3, 1, 2}, {6, 5, 4, 2, 1}},
			[]int{0, 0},
		},
	}

	for _, s := range suites {
		ret := findPeakGrid(s.mat)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
