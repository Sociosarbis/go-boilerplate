package pairs

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	low  int
	high int
	ret  int
}

func TestPairs(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 4, 2, 7},
			2,
			6,
			6,
		},
		{
			[]int{9, 8, 4, 2, 1},
			5,
			14,
			8,
		},
	}

	for _, s := range suites {
		ret := countPairs(s.nums, s.low, s.high)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
