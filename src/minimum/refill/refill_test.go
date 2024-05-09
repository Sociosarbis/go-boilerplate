package refill

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	plants    []int
	capacityA int
	capacityB int
	ret       int
}

func TestRefill(t *testing.T) {
	suites := []suite{
		{
			[]int{2, 2, 3, 3},
			5,
			5,
			1,
		},
		{
			[]int{2, 2, 3, 3},
			3,
			4,
			2,
		},
		{
			[]int{5},
			10,
			8,
			0,
		},
	}

	for _, s := range suites {
		ret := minimumRefill(s.plants, s.capacityA, s.capacityB)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
