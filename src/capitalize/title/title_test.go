package title

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	title string
	ret   string
}

func TestTitle(t *testing.T) {
	suites := []suite{
		{
			"capiTalIze tHe titLe",
			"Capitalize The Title",
		},
		{
			"First leTTeR of EACH Word",
			"First Letter of Each Word",
		},
		{
			"i lOve leetcode",
			"i Love Leetcode",
		},
	}

	for _, s := range suites {
		ret := capitalizeTitle(s.title)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
