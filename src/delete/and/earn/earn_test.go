package earn

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	ret  int
}

func TestEarn(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 4, 2},
			6,
		},
		{
			[]int{2, 2, 3, 3, 3, 4},
			9,
		},
		{
			[]int{8, 7, 3, 8, 1, 4, 10, 10, 10, 2},
			52,
		},
	}
	for _, s := range suites {
		ret := deleteAndEarn(s.nums)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
