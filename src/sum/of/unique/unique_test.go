package unique

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestUnique(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3, 2},
			4,
		},
		{
			[]int{1, 2, 3, 2},
			4,
		},
		{
			[]int{1, 2, 3, 4, 5},
			15,
		},
	}

	for _, su := range suites {
		ret := sumOfUnique(su.nums)
		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
