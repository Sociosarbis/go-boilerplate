package integers

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestIntegers(t *testing.T) {
	suites := []suite{
		{
			5,
			5,
		},
		{
			1,
			2,
		},
		{
			2,
			3,
		},
		{
			3,
			3,
		},
		{
			45,
			21,
		},
	}

	for _, s := range suites {
		ret := findIntegers(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
