package acronym

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	words []string
	s     string
	ret   bool
}

func TestAcronym(t *testing.T) {
	suites := []suite{
		{
			[]string{"alice", "bob", "charlie"},
			"abc",
			true,
		},
		{
			[]string{"an", "apple"},
			"a",
			false,
		},
		{
			[]string{"never", "gonna", "give", "up", "on", "you"},
			"ngguoy",
			true,
		},
	}

	for _, s := range suites {
		ret := isAcronym(s.words, s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
