package sorted

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words []string
	order string
	ret   bool
}

func TestSorted(t *testing.T) {
	suites := []suite{
		{
			[]string{"hello", "leetcode"},
			"hlabcdefgijkmnopqrstuvwxyz",
			true,
		},
		{
			[]string{"word", "world", "row"},
			"worldabcefghijkmnpqstuvxyz",
			false,
		},
		{
			[]string{"apple", "app"},
			"abcdefghijklmnopqrstuvwxyz",
			false,
		},
		{
			[]string{"ubg", "kwh"},
			"qcipyamwvdjtesbghlorufnkzx",
			true,
		},
	}

	for _, s := range suites {
		ret := isAlienSorted(s.words, s.order)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
