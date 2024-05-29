package length

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestLength(t *testing.T) {
	suites := []suite{
		{
			"aaaa",
			2,
		},
		{
			"abcdef",
			-1,
		},
		{
			"abcaba",
			1,
		},
		{
			"cccerrrecdcdccedecdcccddeeeddcdcddedccdceeedccecde",
			2,
		},
	}

	for _, s := range suites {
		ret := maximumLength(s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
