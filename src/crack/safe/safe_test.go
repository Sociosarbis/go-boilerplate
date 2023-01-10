package safe

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	k   int
	ret string
}

func TestSafe(t *testing.T) {
	suites := []suite{
		{
			1,
			2,
			"01",
		},
		{
			2,
			2,
			"01100",
		},
	}

	for _, s := range suites {
		ret := crackSafe(s.n, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
