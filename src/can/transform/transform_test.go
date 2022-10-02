package transform

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	start string
	end   string
	ret   bool
}

func TestTransform(t *testing.T) {
	suites := []suite{
		{
			"RXXLRXRXL",
			"XRLXXRRLX",
			true,
		},
		{
			"X",
			"L",
			false,
		},
	}

	for _, s := range suites {
		ret := canTransform(s.start, s.end)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
