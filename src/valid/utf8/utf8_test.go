package utf8

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	data []int
	ret  bool
}

func TestValidUtf8(t *testing.T) {
	suites := []suite{
		{
			[]int{197, 130, 1},
			true,
		},
		{
			[]int{235, 140, 4},
			false,
		},
		{
			[]int{230, 136, 145},
			true,
		},
	}

	for _, su := range suites {
		ret := validUtf8(su.data)
		if su.ret != ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
