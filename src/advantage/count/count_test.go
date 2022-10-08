package count

import (
	"reflect"
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums1 []int
	nums2 []int
	ret   []int
}

func TestCount(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 7, 11, 15},
			[]int{1, 10, 4, 11},
			[]int{2, 11, 7, 15},
		},
		{
			[]int{12, 24, 8, 32},
			[]int{13, 25, 32, 11},
			[]int{24, 32, 8, 12},
		},
	}
	for _, s := range suites {
		ret := advantageCount(s.nums1, s.nums2)
		if !reflect.DeepEqual(s.ret, ret) {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
