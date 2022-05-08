package duplicates

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  []int
}

func TestDuplicates(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{3, 2},
		},
		{
			[]int{1, 1, 2},
			[]int{1},
		},
		{
			[]int{1},
			[]int{},
		},
	}
	for _, s := range suites {
		ret := findDuplicates(s.nums)

		if !reflect.DeepEqual(ret, s.ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
