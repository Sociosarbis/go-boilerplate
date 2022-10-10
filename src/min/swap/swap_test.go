package swap

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums1 []int
	nums2 []int
	ret   int
}

func TestSwap(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 3, 5, 4},
			[]int{1, 2, 3, 7},
			1,
		},
		{
			[]int{0, 3, 5, 8, 9},
			[]int{2, 1, 4, 6, 9},
			1,
		},
	}

	for _, s := range suites {
		ret := minSwap(s.nums1, s.nums2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
