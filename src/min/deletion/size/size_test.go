package size

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	strs []string
	ret  int
}

func TestMinDeletionSize(t *testing.T) {
	suites := []suite{
		{
			[]string{"cba", "daf", "ghi"},
			1,
		},
		{
			[]string{"a", "b"},
			0,
		},
		{
			[]string{"zyx", "wvu", "tsr"},
			3,
		},
	}

	for _, s := range suites {
		ret := minDeletionSize(s.strs)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
