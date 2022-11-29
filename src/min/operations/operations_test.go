package operations

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestOperations(t *testing.T) {
	suites := []suite{
		{
			"0100",
			1,
		},
		{
			"10",
			0,
		},
	}
	for _, s := range suites {
		ret := minOperations(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
