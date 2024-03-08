package sum

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n      int
	target int
	ret    int
}

func TestSum(t *testing.T) {
	suites := []suite{{
		2,
		3,
		4,
	}, {
		3,
		3,
		8,
	}, {
		1,
		1,
		1,
	}, {
		1000000000,
		1000000000,
		750000042,
	}}

	for _, s := range suites {
		ret := minimumPossibleSum(s.n, s.target)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
