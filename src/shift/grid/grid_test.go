package grid

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	grid [][]int
	k    int
	ret  [][]int
}

func TestGrid(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			1,
			[][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			4,
			[][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			9,
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}

	for _, s := range suites {
		ret := shiftGrid(s.grid, s.k)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
