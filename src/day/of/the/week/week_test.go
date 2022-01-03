package week

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	day   int
	month int
	year  int
	ret   string
}

func TestWeek(t *testing.T) {
	suites := []suite{
		{
			31,
			8,
			2019,
			"Saturday",
		},
		{
			18,
			7,
			1999,
			"Sunday",
		},
		{
			15,
			8,
			1993,
			"Sunday",
		},
	}

	for _, su := range suites {
		ret := dayOfTheWeek(su.day, su.month, su.year)
		if ret != su.ret {
			t.Errorf(constant.ErrTemplateStr, su.ret, ret)
		}
	}
}
