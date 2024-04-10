package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	binary string
	ret    string
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			"000110",
			"111011",
		},
		{
			"01",
			"01",
		},
	}

	for _, s := range suites {
		ret := maximumBinaryString(s.binary)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
