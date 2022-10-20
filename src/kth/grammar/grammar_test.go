package grammar

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	k   int
	ret int
}

func TestGrammar(t *testing.T) {
	suites := []suite{
		{
			1,
			1,
			0,
		},
		{
			2,
			1,
			0,
		},
		{
			2,
			2,
			1,
		},
	}

	for _, s := range suites {
		ret := kthGrammar(s.n, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
