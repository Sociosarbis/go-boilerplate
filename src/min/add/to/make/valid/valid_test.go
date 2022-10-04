package valid

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestValid(t *testing.T) {
	suites := []suite{
		{
			"())",
			1,
		},
		{
			"(((",
			3,
		},
	}

	for _, s := range suites {
		ret := minAddToMakeValid(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.s, s.ret)
		}
	}
}
