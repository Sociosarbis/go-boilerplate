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
			[]int{1},
			1,
			1,
		},
		{
			[]int{1, 2},
			4,
			-1,
		},
		{
			[]int{2, -1, 2},
			3,
			3,
		},
	}

	for _, s := range suites {
		ret := shortestSubarray(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
