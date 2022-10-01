package number

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	number string
	ret    string
}

func TestNumber(t *testing.T) {
	suites := []suite{
		{
			"1-23-45 6",
			"123-456",
		},
		{
			"123 4-567",
			"123-45-67",
		},
		{
			"123 4-5678",
			"123-456-78",
		},
	}

	for _, s := range suites {
		ret := reformatNumber(s.number)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
