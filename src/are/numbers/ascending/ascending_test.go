package ascending

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret bool
}

func TestAscending(t *testing.T) {
	suites := []suite{
		{
			"1 box has 3 blue 4 red 6 green and 12 yellow marbles",
			true,
		},
		{
			"hello world 5 x 5",
			false,
		},
		{
			"sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s",
			false,
		},
	}

	for _, s := range suites {
		ret := areNumbersAscending(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
