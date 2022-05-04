package winner

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	k   int
	ret int
}

func TestWinner(t *testing.T) {
	suites := []suite{
		{
			5,
			2,
			3,
		},
		{
			6,
			5,
			1,
		},
	}

	for _, s := range suites {
		ret := findTheWinner(s.n, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
