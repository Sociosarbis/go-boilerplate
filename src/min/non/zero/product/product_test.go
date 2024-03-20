package product

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	p   int
	ret int
}

func TestProduct(t *testing.T) {
	suites := []suite{
		{
			1,
			1,
		},
		{
			2,
			6,
		},
		{
			3,
			1512,
		},
	}

	for _, s := range suites {
		ret := minNonZeroProduct(s.p)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
