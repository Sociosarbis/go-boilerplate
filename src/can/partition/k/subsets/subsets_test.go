package subsets

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  bool
}

func TestSubsets(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 3, 2, 3, 5, 2, 1},
			4,
			true,
		},
		{
			[]int{1, 2, 3, 4},
			3,
			false,
		},
	}

	for _, s := range suites {
		ret := canPartitionKSubsets(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
