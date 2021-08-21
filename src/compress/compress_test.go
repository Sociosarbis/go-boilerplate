package compress

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	chars []byte
	ret   int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			6,
		},
		{
			[]byte{'a'},
			1,
		},
		{
			[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			4,
		},
	}

	for _, s := range suites {
		ret := compress(s.chars)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
