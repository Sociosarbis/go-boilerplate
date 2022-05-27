package closest

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words []string
	word1 string
	word2 string
	ret   int
}

func TestClosest(t *testing.T) {
	suites := []suite{
		{
			[]string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"},
			"a",
			"student",
			1,
		},
	}

	for _, s := range suites {
		ret := findClosest(s.words, s.word1, s.word2)

		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
