package alternately

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	word1 string
	word2 string
	ret   string
}

func TestAlternately(t *testing.T) {
	suites := []suite{
		{
			"abc",
			"pqr",
			"apbqcr",
		},
		{
			"ab",
			"pqrs",
			"apbqrs",
		},
		{
			"abcd",
			"pq",
			"apbqcd",
		},
	}

	for _, s := range suites {
		ret := mergeAlternately(s.word1, s.word2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
