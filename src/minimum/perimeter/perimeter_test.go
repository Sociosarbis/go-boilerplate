package perimeter

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	neededApples int64
	ret          int64
}

func TestPerimeter(t *testing.T) {
	suites := []suite{
		{
			1,
			8,
		}, {
			13,
			16,
		},
		{
			1000000000,
			5040,
		},
	}

	for _, s := range suites {
		ret := minimumPerimeter(s.neededApples)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
