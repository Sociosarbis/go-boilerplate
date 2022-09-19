package sort

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestSort(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 2, 2, 2, 3},
			[]int{3, 1, 1, 2, 2, 2},
		},
		{
			[]int{2, 3, 1, 3, 2},
			[]int{1, 3, 3, 2, 2},
		},
		{
			[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			[]int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}

	for _, s := range suites {
		ret := frequencySort(s.nums)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
