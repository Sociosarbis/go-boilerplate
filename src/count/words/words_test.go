package words

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words1 []string
	words2 []string
	ret    int
}

func TestWords(t *testing.T) {
	suites := []suite{
		{
			[]string{"leetcode", "is", "amazing", "as", "is"},
			[]string{"amazing", "leetcode", "is"},
			2,
		},
		{
			[]string{"b", "bb", "bbb"},
			[]string{"a", "aa", "aaa"},
			0,
		},
		{
			[]string{"a", "ab"},
			[]string{"a", "a", "a", "ab"},
			1,
		},
	}

	for _, s := range suites {
		ret := countWords(s.words1, s.words2)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
