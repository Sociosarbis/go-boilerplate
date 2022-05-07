package mutation

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	start string
	end   string
	bank  []string
	ret   int
}

func TestMutation(t *testing.T) {
	suites := []suite{
		{
			"AACCGGTT",
			"AACCGGTA",
			[]string{"AACCGGTA"},
			1,
		},
		{
			"AACCGGTT",
			"AAACGGTA",
			[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			2,
		},
		{
			"AAAAACCC",
			"AACCCCCC",
			[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			3,
		},
	}

	for _, s := range suites {
		ret := minMutation(s.start, s.end, s.bank)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
