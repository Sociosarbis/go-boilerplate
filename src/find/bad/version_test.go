package findBadVersion

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	n   int
	bad int
}

func TestResults(t *testing.T) {
	suites := []suite{
		{
			5,
			4,
		},
		{
			1,
			1,
		},
	}

	for _, s := range suites {
		badVersion = s.bad
		ret := firstBadVersion(s.n)
		if ret != badVersion {
			t.Errorf(constant.ErrTemplateStr, badVersion, ret)
		}
	}
}
