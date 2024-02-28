package increments

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n    int
	cost []int
	ret  int
}

func TestIncrements(t *testing.T) {
	suites := []suite{
		{
			7,
			[]int{1, 5, 2, 2, 3, 3, 1},
			6,
		},
		{
			3,
			[]int{5, 3, 3},
			0,
		},
		{
			31,
			[]int{9081, 3243, 2264, 8676, 8749, 1573, 483, 3339, 3037, 1272, 6398, 4801, 5149, 5583, 8338, 6709, 5026, 5608, 9387, 4730, 9548, 9856, 2868, 6805, 6805, 8847, 5086, 7393, 6303, 6045, 7778},
			51508,
		},
	}

	for _, s := range suites {
		ret := minIncrements(s.n, s.cost)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
