package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			1,
			22,
		},
		{
			1000,
			1333,
		},
		{
			3000,
			3133,
		},
	}

	for _, s := range suites {
		ret := nextBeautifulNumber(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}

}
