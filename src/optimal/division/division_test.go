package division

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  string
}

func TestDivision(t *testing.T) {
	suites := []suite{
		{
			[]int{1000, 100, 10, 2},
			"1000/(100/10/2)",
		},
		{
			[]int{2, 3, 4},
			"2/(3/4)",
		},
		{
			[]int{2},
			"2",
		},
	}

	for _, s := range suites {
		ret := optimalDivision(s.nums)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
