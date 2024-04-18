package array

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	changed []int
	ret     []int
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 3, 4, 2, 6, 8},
			[]int{1, 3, 4},
		},
		{
			[]int{6, 3, 0, 1},
			[]int{},
		},
		{
			[]int{1},
			[]int{},
		},
	}

	for _, s := range suites {
		ret := findOriginalArray(s.changed)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
