package difference

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret [][]int
}

func TestDifferecne(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 2, 1, 3},
			[][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			[]int{1, 3, 6, 10, 15},
			[][]int{{1, 3}},
		},
		{
			[]int{3, 8, -10, 23, 19, -4, -14, 27},
			[][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for _, s := range suites {
		ret := minimumAbsDifference(s.arr)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
