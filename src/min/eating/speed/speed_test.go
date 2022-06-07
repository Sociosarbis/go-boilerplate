package speed

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	piles []int
	h     int
	ret   int
}

func TestSpeed(t *testing.T) {
	suites := []suite{
		{
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			[]int{30, 11, 23, 4, 20},
			5,
			30,
		},
		{
			[]int{30, 11, 23, 4, 20},
			6,
			23,
		},
	}

	for _, s := range suites {
		ret := minEatingSpeed(s.piles, s.h)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
