package formatting

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	k   int
	ret string
}

func TestFormatting(t *testing.T) {
	suites := []suite{
		{
			"5F3Z-2e-9-w",
			4,
			"5F3Z-2E9W",
		},
		{
			"2-5g-3-J",
			2,
			"2-5G-3J",
		},
	}

	for _, s := range suites {
		ret := licenseKeyFormatting(s.s, s.k)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
