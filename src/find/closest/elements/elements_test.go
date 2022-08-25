package elements

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	k   int
	x   int
	ret []int
}

func TestElements(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3, 4, 5},
			4,
			3,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
			-1,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8},
			2,
			2,
			[]int{1, 3},
		},
	}
	for _, s := range suites {
		ret := findClosestElements(s.arr, s.k, s.x)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
