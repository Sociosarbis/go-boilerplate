package addr

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	address string
	ret     string
}

func TestAddr(t *testing.T) {
	suites := []suite{
		{
			"1.1.1.1",
			"1[.]1[.]1[.]1",
		},
		{
			"255.100.50.0",
			"255[.]100[.]50[.]0",
		},
	}

	for _, s := range suites {
		ret := defangIPaddr(s.address)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
