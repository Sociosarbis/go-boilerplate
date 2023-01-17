package pairs

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestPairs(t *testing.T) {
	suites := []suite{
		{
			[]int{42, 11, 1, 97},
			2,
		},
		{
			[]int{13, 10, 35, 24, 76},
			4,
		},
	}

	for _, s := range suites {
		ret := countNicePairs(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
