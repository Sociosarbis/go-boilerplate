package highest

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestHighest(t *testing.T) {
	suites := []suite{
		{
			"dfa12321afd",
			2,
		},
		{
			"abc111",
			-1,
		},
	}

	for _, s := range suites {
		ret := secondHighest(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
