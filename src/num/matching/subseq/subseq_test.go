package subseq

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s     string
	words []string
	ret   int
}

func TestSubseq(t *testing.T) {
	suites := []suite{
		{
			"abcde",
			[]string{"a", "bb", "acd", "ace"},
			3,
		},
		{
			"dsahjpjauf",
			[]string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"},
			2,
		},
	}

	for _, s := range suites {
		ret := numMatchingSubseq(s.s, s.words)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
