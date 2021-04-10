package ugly

import (
	"testing"
)

type suite struct {
	n   int
	ret bool
}

func TestIsUgly(t *testing.T) {
	suites := []suite{
		{
			6,
			true,
		},
		{
			8,
			true,
		},
		{
			14,
			false,
		},
	}

	for _, s := range suites {
		ret := isUgly(s.n)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
