package special

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	mat [][]int
	ret int
}

func TestSpecial(t *testing.T) {
	suites := []suite{
		{
			[][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}},
			1,
		},
		{
			[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			3,
		},
	}
	for _, s := range suites {
		ret := numSpecial(s.mat)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
