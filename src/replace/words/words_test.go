package words

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	dictionary []string
	sentence   string
	ret        string
}

func TestWords(t *testing.T) {
	suites := []suite{
		{
			[]string{"cat", "bat", "rat"},
			"the cattle was rattled by the battery",
			"the cat was rat by the bat",
		},
		{
			[]string{"a", "b", "c"},
			"aadsfasf absbs bbab cadsfafs",
			"a a b c",
		},
	}

	for _, s := range suites {
		ret := replaceWords(s.dictionary, s.sentence)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
