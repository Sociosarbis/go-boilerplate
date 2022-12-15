package lucky

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	k   int
	ret int
}

func TestLucky(t *testing.T) {
	suites := []suite{
		{
			"iiii",
			1,
			36,
		},
		{
			"leetcode",
			2,
			6,
		},
		{
			"zbax",
			2,
			8,
		},
	}

	for _, s := range suites {
		ret := getLucky(s.s, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
