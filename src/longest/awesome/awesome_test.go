package awesome

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	s   string
	ret int
}

func TestAwesome(t *testing.T) {
	suites := []suite{
		{
			"3242415",
			5,
		},
		{
			"12345678",
			1,
		},
		{
			"213123",
			6,
		},
		{
			"76263",
			3,
		},
		{
			"185801630663498",
			5,
		},
		{
			"54963750294999325",
			5,
		},
	}

	for _, s := range suites {
		ret := longestAwesome(s.s)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
