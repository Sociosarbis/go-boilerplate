package balloons

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	text string
	ret  int
}

func TestBalloons(t *testing.T) {
	suites := []suite{
		{
			"nlaebolko",
			1,
		},
		{
			"loonbalxballpoon",
			2,
		},
		{
			"leetcode",
			0,
		},
	}

	for _, s := range suites {
		ret := maxNumberOfBalloons(s.text)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
