package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	k   int
	ret int
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			5,
			9,
		},
	}

	for _, s := range suites {
		ret := getKthMagicNumber(s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
