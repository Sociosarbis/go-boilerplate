package words

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s     string
	words []string
	ret   int
}

func TestWords(t *testing.T) {
	suites := []suite{
		{
			"heeellooo",
			[]string{"hello", "hi", "helo"},
			1,
		},
		{
			"zzzzzyyyyy",
			[]string{"zzyy", "zy", "zyy"},
			3,
		},
	}

	for _, s := range suites {
		ret := expressiveWords(s.s, s.words)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
