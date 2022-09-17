package characters

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestCharacters(t *testing.T) {
	suites := []suite{
		{
			"aa",
			0,
		},
		{
			"abca",
			2,
		},
		{
			"cbzxy",
			-1,
		},
	}
	for _, s := range suites {
		ret := maxLengthBetweenEqualCharacters(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
