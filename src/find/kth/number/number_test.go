package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	k   int
	ret int
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			13,
			2,
			10,
		},
		{
			1,
			1,
			1,
		},
		{
			1000,
			1000,
			999,
		},
	}

	for _, s := range suites {
		ret := findKthNumber(s.n, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
