package score2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums []int
	x    int
	ret  int64
}

func TestScore(t *testing.T) {
	suites := []suite{
		{
			[]int{9, 58, 17, 54, 91, 90, 32, 6, 13, 67, 24, 80, 8, 56, 29, 66, 85, 38, 45, 13, 20, 73, 16, 98, 28, 56, 23, 2, 47, 85, 11, 97, 72, 2, 28, 52, 33},
			90,
			886,
		},
	}

	for _, s := range suites {
		ret := maxScore(s.nums, s.x)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
