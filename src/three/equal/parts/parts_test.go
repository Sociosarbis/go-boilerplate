package parts

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	arr []int
	ret []int
}

func TestParts(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 0, 1, 0, 1},
			[]int{0, 3},
		},
		{
			[]int{1, 1, 0, 1, 1},
			[]int{-1, -1},
		},
		{
			[]int{1, 1, 0, 0, 1},
			[]int{0, 2},
		},
		{
			[]int{1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0},
			[]int{15, 32},
		},
	}

	for _, s := range suites {
		ret := threeEqualParts(s.arr)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
