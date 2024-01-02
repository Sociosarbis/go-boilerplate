package repetitions

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s1  string
	n1  int
	s2  string
	n2  int
	ret int
}

func TestRepetitions(t *testing.T) {
	suites := []suite{
		{
			"acb",
			4,
			"ab",
			2,
			2,
		},
		{
			"acb",
			1,
			"acb",
			1,
			1,
		},
	}

	for _, s := range suites {
		ret := getMaxRepetitions(s.s1, s.n1, s.s2, s.n2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
