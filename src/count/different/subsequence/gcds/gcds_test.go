package gcds

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestGCDs(t *testing.T) {
	suites := []suite{
		{
			[]int{6, 10, 3},
			5,
		},
		{
			[]int{5, 15, 40, 5, 6},
			7,
		},
	}

	for _, s := range suites {
		ret := countDifferentSubsequenceGCDs(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
