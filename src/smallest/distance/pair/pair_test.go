package pair

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	k    int
	ret  int
}

func TestPair(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 3, 1},
			1,
			0,
		},
		{
			[]int{1, 1, 1},
			2,
			0,
		},
		{
			[]int{1, 6, 1},
			3,
			5,
		},
	}

	for _, s := range suites {
		ret := smallestDistancePair(s.nums, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
