package subarray

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestSubarray(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 3, 2, 3, 1, 3},
			3,
			3,
		},
		{
			[]int{1, 1, 2, 2, 1, 1},
			2,
			4,
		},
		{
			[]int{2, 4, 2, 6, 2, 9, 5, 7, 8, 6, 5, 6, 2, 3, 8, 6, 2, 4, 1, 6, 2, 7, 6, 8, 8, 3, 6, 5, 10, 3, 5, 7},
			8,
			4,
		},
	}

	for _, s := range suites {
		ret := longestEqualSubarray(s.nums, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
