package count

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	num1    string
	num2    string
	min_num int
	max_num int
	ret     int
}

func TestCount(t *testing.T) {
	suites := []suite{
		{
			"1",
			"12",
			1,
			8,
			11,
		},
		{
			"1",
			"5",
			1,
			5,
			5,
		},
		{
			"4179205230",
			"7748704426",
			8,
			46,
			883045899,
		},
	}

	for _, s := range suites {
		ret := count(s.num1, s.num2, s.min_num, s.max_num)

		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
