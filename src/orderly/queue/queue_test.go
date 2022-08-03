package queue

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	k   int
	ret string
}

func TestOrderlyQueue(t *testing.T) {
	suites := []suite{
		{
			"cba",
			1,
			"acb",
		},
		{
			"baaca",
			3,
			"aaabc",
		},
	}

	for _, s := range suites {
		ret := orderlyQueue(s.s, s.k)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
