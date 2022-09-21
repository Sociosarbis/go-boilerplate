package similarity

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s1  string
	s2  string
	ret int
}

func TestSimilarity(t *testing.T) {
	suites := []suite{
		{
			"ab",
			"ba",
			1,
		},
		{
			"abc",
			"bca",
			2,
		},
	}
	for _, s := range suites {
		ret := kSimilarity(s.s1, s.s2)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
