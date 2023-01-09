package permutation

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestPermutation(t *testing.T) {
	suites := []suite{
		{
			2,
			1,
		},
		{
			4,
			2,
		},
		{
			6,
			4,
		},
	}

	for _, s := range suites {
		ret := reinitializePermutation(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
