package substrings

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	t   string
	ret int
}

func TestSubstrings(t *testing.T) {
	suites := []suite{
		{
			"aba",
			"baba",
			6,
		},
		{
			"ab",
			"bb",
			3,
		},
	}

	for _, s := range suites {
		ret := countSubstrings(s.s, s.t)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
