package operation3

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums1 []int
	nums2 []int
	ret   int
}

func TestOperation(t *testing.T) {
	suites := []suite{
		/*{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 1, 2, 2, 2, 2},
			3,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			[]int{6},
			-1,
		},
		{
			[]int{6, 6},
			[]int{1},
			3,
		},*/
		{
			[]int{5, 2, 1, 5, 2, 2, 2, 2, 4, 3, 3, 5},
			[]int{1, 4, 5, 5, 6, 3, 1, 3, 3},
			1,
		},
	}

	for _, s := range suites {
		ret := minOperations(s.nums1, s.nums2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
