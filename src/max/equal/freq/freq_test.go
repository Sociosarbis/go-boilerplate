package freq

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestFreq(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 2, 1, 1, 5, 3, 3, 5},
			7,
		},
		{
			[]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5},
			13,
		},
	}

	for _, s := range suites {
		ret := maxEqualFreq(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
