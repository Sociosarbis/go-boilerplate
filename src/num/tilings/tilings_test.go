package tilings

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestTilings(t *testing.T) {
	suites := []suite{
		{
			3,
			5,
		},
		{
			1,
			1,
		},
		{
			4,
			11,
		},
	}

	for _, s := range suites {
		ret := numTilings(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
