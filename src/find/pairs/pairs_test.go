package pairs

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestPairs(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 1, 4, 1, 5},
			2,
			2,
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			4,
		},
		{
			[]int{1, 3, 1, 5, 4},
			0,
			1,
		},
		{
			[]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3},
			3,
			2,
		},
	}

	for _, s := range suites {
		ret := findPairs(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
