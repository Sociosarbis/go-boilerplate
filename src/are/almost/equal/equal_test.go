package equal

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s1  string
	s2  string
	ret bool
}

func TestEqual(t *testing.T) {
	suites := []suite{
		{
			"bank",
			"kanb",
			true,
		},
		{
			"attack",
			"defend",
			false,
		},
		{
			"kelb",
			"kelb",
			true,
		},
	}

	for _, s := range suites {
		ret := areAlmostEqual(s.s1, s.s2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
