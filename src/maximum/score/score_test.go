package score

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	a   int
	b   int
	c   int
	ret int
}

func TestScore(t *testing.T) {
	suites := []suite{
		{
			2, 4, 6, 6,
		},
		{
			4, 4, 6, 7,
		},
		{
			1, 8, 8, 8,
		},
	}

	for _, s := range suites {
		ret := maximumScore(s.a, s.b, s.c)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
