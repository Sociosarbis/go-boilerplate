package isomophic

import (
	"testing"
)

type suite struct {
	s   string
	t   string
	ret bool
}

func TestIsIsomorphic(t *testing.T) {
	suites := []suite{
		{
			"egg",
			"add",
			true,
		},
		{
			"foo",
			"bar",
			false,
		},
		{
			"paper",
			"title",
			true,
		},
	}

	for _, s := range suites {
		ret := isIsomorphic(s.s, s.t)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
