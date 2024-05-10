package devices

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	batteryPercentages []int
	ret                int
}

func TestDevices(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 2, 1, 3},
			3,
		},
		{
			[]int{0, 1, 2},
			2,
		},
	}

	for _, s := range suites {
		ret := countTestedDevices(s.batteryPercentages)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
