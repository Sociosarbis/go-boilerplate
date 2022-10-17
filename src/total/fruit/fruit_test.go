package fruit

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	fruits []int
	ret    int
}

func TestFruit(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 1},
			3,
		},
		{
			[]int{0, 1, 2, 2},
			3,
		},
		{
			[]int{1, 2, 3, 2, 2},
			4,
		},
		{
			[]int{0, 1, 1, 4, 3},
			3,
		},
	}

	for _, s := range suites {
		ret := totalFruit(s.fruits)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
