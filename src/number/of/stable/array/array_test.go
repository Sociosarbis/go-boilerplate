package array

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	zero  int
	one   int
	limit int
	ret   int
}

func TestArray(t *testing.T) {
	suites := []suite{
		{
			1,
			1,
			2,
			2,
		},
		{
			1,
			2,
			1,
			1,
		},
		{
			1,
			2,
			2,
			3,
		},
		{
			3,
			3,
			2,
			14,
		},
	}

	for _, s := range suites {
		ret := numberOfStableArrays(s.zero, s.one, s.limit)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
