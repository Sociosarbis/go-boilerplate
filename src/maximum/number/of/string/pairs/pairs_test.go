package pairs

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words []string
	ret   int
}

func TestPairs(t *testing.T) {
	suites := []suite{
		{
			[]string{"cd", "ac", "dc", "ca", "zz"},
			2,
		},
		{
			[]string{"ab", "ba", "cc"},
			1,
		},
	}

	for _, s := range suites {
		ret := maximumNumberOfStringPairs(s.words)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
