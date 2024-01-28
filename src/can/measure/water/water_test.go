package water

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	jug1Capacity   int
	jug2Capacity   int
	targetCapacity int
	ret            bool
}

func TestWater(t *testing.T) {
	suites := []suite{
		{
			3,
			5,
			4,
			true,
		},
		{
			2,
			6,
			5,
			false,
		},
		{
			1,
			2,
			3,
			true,
		},
	}

	for _, s := range suites {
		ret := canMeasureWater(s.jug1Capacity, s.jug2Capacity, s.targetCapacity)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
