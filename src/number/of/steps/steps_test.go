package steps

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	num int
	ret int
}

func TestSteps(t *testing.T) {
	suites := []suite{
		{
			14,
			6,
		},
		{
			8,
			4,
		},
	}
	for _, su := range suites {
		ret := numberOfSteps(su.num)
		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, ret, su.ret)
		}
	}
}
