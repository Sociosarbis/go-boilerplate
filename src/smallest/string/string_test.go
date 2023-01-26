package string

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	k   int
	ret string
}

func TestString(t *testing.T) {
	suites := []suite{
		{
			3,
			27,
			"aay",
		},
		{
			5,
			73,
			"aaszz",
		},
	}

	for _, s := range suites {
		ret := getSmallestString(s.n, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
