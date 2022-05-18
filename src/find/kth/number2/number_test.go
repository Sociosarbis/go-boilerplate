package number2

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	m   int
	n   int
	k   int
	ret int
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			3,
			3,
			5,
			3,
		},
		{
			2,
			3,
			6,
			6,
		},
	}

	for _, s := range suites {
		ret := findKthNumber(s.m, s.n, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
