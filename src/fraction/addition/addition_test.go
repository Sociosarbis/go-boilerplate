package addition

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	expression string
	ret        string
}

func TestAddition(t *testing.T) {
	suites := []suite{
		/*{
			"-1/2+1/2",
			"0/1",
		},
		{
			"-1/2+1/2+1/3",
			"1/3",
		},*/
		{
			"1/3-1/2",
			"-1/6",
		},
	}

	for _, s := range suites {
		ret := fractionAddition(s.expression)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
