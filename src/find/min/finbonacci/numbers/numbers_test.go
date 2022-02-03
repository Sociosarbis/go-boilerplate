package numbers

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	k   int
	ret int
}

func TestNumbers(t *testing.T) {
	suites := []suite{
		{
			7,
			2,
		},
		{
			10,
			2,
		},
		{
			19,
			3,
		},
	}

	for _, su := range suites {
		ret := findMinFibonacciNumbers(su.k)

		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
