package merge

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	word1 string
	word2 string
	ret   string
}

func TestMerge(t *testing.T) {
	suites := []suite{
		{
			"cabaa",
			"bcaaa",
			"cbcabaaaaa",
		},
		{
			"abcabc",
			"abdcaba",
			"abdcabcabcaba",
		},
	}

	for _, s := range suites {
		ret := largestMerge(s.word1, s.word2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
