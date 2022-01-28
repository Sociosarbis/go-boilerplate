package characters

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	properties [][]int
	ret        int
}

func TestCharacters(t *testing.T) {
	suites := []suite{
		{
			[][]int{{5, 5}, {6, 3}, {3, 6}},
			0,
		},
		{
			[][]int{{2, 2}, {3, 3}},
			1,
		},
		{
			[][]int{{1, 5}, {10, 4}, {4, 3}},
			1,
		},
	}

	for _, su := range suites {
		ret := numberOfWeakCharacters(su.properties)
		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
