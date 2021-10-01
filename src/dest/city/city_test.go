package city

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	paths [][]string
	ret   string
}

func TestCitys(t *testing.T) {
	suites := []suite{
		{
			[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			"Sao Paulo",
		},
		{
			[][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			"A",
		},
	}

	for _, s := range suites {
		ret := destCity(s.paths)
		if s.ret != ret {
			t.Errorf(constant.ErrTemplateStr, ret, s.ret)
		}
	}
}
