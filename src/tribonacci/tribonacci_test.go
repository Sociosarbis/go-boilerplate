package tribonacci

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	ret int
}

func TestTribonacci(t *testing.T) {
	suites := []suite{
		{
			4,
			4,
		},
		{
			25,
			1389537,
		},
	}

	for _, s := range suites {
		ret := tribonacci(s.n)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
