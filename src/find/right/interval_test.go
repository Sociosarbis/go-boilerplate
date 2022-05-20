package interval

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	intervals [][]int
	ret       []int
}

func TestInterval(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 2}},
			[]int{-1},
		},
		{
			[][]int{{3, 4}, {2, 3}, {1, 2}},
			[]int{-1, 0, 1},
		},
		{
			[][]int{{1, 4}, {2, 3}, {3, 4}},
			[]int{-1, 2, -1},
		},
	}

	for _, s := range suites {
		ret := findRightInterval(s.intervals)

		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
