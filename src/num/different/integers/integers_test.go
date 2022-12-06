package integers

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	word string
	ret  int
}

func TestIntegers(t *testing.T) {
	suites := []suite{
		{
			"a123bc34d8ef34",
			3,
		},
		{
			"leet1234code234",
			2,
		},
		{
			"a1b01c001",
			1,
		},
	}

	for _, s := range suites {
		ret := numDifferentIntegers(s.word)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
