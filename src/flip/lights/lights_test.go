package lights

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n       int
	presses int
	ret     int
}

func TestLights(t *testing.T) {
	suites := []suite{
		{
			1,
			1,
			2,
		},
		{
			2,
			1,
			3,
		},
		{
			3,
			1,
			4,
		},
	}

	for _, s := range suites {
		ret := flipLights(s.n, s.presses)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
