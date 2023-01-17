package similar

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	sentence1 string
	sentence2 string
	ret       bool
}

func TestSimilar(t *testing.T) {
	suites := []suite{
		{
			"My name is Haley",
			"My Haley",
			true,
		},
		{
			"of",
			"A lot of words",
			false,
		},
		{
			"Eating right now",
			"Eating",
			true,
		},
	}

	for _, s := range suites {
		ret := areSentencesSimilar(s.sentence1, s.sentence2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
