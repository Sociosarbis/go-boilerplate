package elements

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	nums  []int
	limit int
	goal  int
	ret   int
}

func TestElements(t *testing.T) {
	suites := []suite{
		{
			[]int{1, -1, 1},
			3,
			-4,
			2,
		},
		{
			[]int{1, -10, 9, 1},
			100,
			0,
			1,
		},
	}

	for _, s := range suites {
		ret := minElements(s.nums, s.limit, s.goal)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
