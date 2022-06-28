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
			[]int{1, 5, 1, 1, 6, 4},
			[]int{1, 5, 1, 6, 1, 4},
		},
		{
			[]int{1, 3, 2, 2, 3, 1},
			[]int{1, 3, 2, 3, 1, 2},
		},
	}

	for _, s := range suites {
		wiggleSort(s.nums)
		if !reflect.DeepEqual(s.ret, s.nums) {
			t.Errorf(constant.ErrTemplateStr, s.ret, s.nums)
		}
	}
}
