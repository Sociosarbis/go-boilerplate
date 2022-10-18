package set

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	digits []string
	n      int
	ret    int
}

func TestSet(t *testing.T) {
	suites := []suite{
		{
			[]string{"1", "3", "5", "7"},
			199,
			36,
		},
		{
			[]string{"1", "4", "9"},
			1000000000,
			29523,
		},
		{
			[]string{"7"},
			8,
			1,
		},
	}

	for _, s := range suites {
		ret := atMostNGivenDigitSet(s.digits, s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
