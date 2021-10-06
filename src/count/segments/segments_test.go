package segments

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestSegments(t *testing.T) {
	suites := []suite{
		{
			"Hello, my name is John",
			5,
		},
		{
			"Hello",
			1,
		},
		{
			"love live! mu'sic forever",
			4,
		},
		{
			"",
			0,
		},
	}

	for _, s := range suites {
		ret := countSegments(s.s)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
