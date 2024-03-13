package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret string
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			"010",
			"001",
		},
		{
			"0101",
			"1001",
		},
	}

	for _, s := range suites {
		ret := maximumOddBinaryNumber(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
