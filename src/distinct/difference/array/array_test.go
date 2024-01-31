package array

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestArray(t *testing.T) {
	suites := []suite{
		{[]int{1, 2, 3, 4, 5},
			[]int{-3, -1, 1, 3, 5}},
		{
			[]int{3, 2, 3, 4, 2},
			[]int{-2, -1, 0, 2, 3},
		},
	}

	for _, s := range suites {
		ret := distinctDifferenceArray(s.nums)
		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
