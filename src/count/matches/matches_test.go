package matches

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	items     [][]string
	ruleKey   string
	ruleValue string
	ret       int
}

func TestMatches(t *testing.T) {
	suites := []suite{
		{
			[][]string{
				{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"},
			},
			"color",
			"silver",
			1,
		},
		{
			[][]string{
				{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"},
			},
			"type",
			"phone",
			2,
		},
	}

	for _, s := range suites {
		ret := countMatches(s.items, s.ruleKey, s.ruleValue)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
