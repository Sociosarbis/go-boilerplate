package length

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestLength(t *testing.T) {
	suites := []suite{
		{
			"ABFCACDB",
			2,
		},
		{
			"ACBBD",
			5,
		},
	}

	for _, s := range suites {
		ret := minLength(s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
