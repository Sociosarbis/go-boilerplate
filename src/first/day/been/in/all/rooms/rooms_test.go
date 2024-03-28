package rooms

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nextVisit []int
	ret       int
}

func TestRooms(t *testing.T) {
	suites := []suite{
		{
			[]int{0, 0},
			2,
		},
		{
			[]int{0, 0, 2},
			6,
		},
		{
			[]int{0, 1, 2, 0},
			6,
		},
	}

	for _, s := range suites {
		ret := firstDayBeenInAllRooms(s.nextVisit)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
