package inclusion

import (
	"testing"
)

type suite struct {
	s1  string
	s2  string
	ret bool
}

func TestInclusion(t *testing.T) {
	suites := []suite{
		{
			"ab",
			"eidbaooo",
			true,
		},
		{
			"ab",
			"eidboaoo",
			false,
		},
	}

	for _, su := range suites {
		ret := checkInclusion(su.s1, su.s2)
		if ret != su.ret {
			t.Errorf("Failed, expected %v but get %v", su.ret, ret)
		}
	}
}
