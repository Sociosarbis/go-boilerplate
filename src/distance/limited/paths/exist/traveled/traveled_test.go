package traveled

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	mainTank       int
	additionalTank int
	ret            int
}

func TestTraveled(t *testing.T) {
	suites := []suite{
		{
			5,
			10,
			60,
		},
		{
			1,
			2,
			10,
		},
	}

	for _, s := range suites {
		ret := distanceTraveled(s.mainTank, s.additionalTank)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
