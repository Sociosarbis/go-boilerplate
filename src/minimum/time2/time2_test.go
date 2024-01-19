package time2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums1 []int
	nums2 []int
	x     int
	ret   int
}

func TestTime(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			4,
			3,
		},
		{
			[]int{1, 2, 3},
			[]int{3, 3, 3},
			4,
			-1,
		},
	}

	for _, s := range suites {
		ret := minimumTime(s.nums1, s.nums2, s.x)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
