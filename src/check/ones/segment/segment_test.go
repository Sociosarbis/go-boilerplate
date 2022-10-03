package segment

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret bool
}

func TestSegment(t *testing.T) {
	suites := []suite{
		{
			"1001",
			false,
		},
		{
			"110",
			true,
		},
	}

	for _, s := range suites {
		ret := checkOnesSegment(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
