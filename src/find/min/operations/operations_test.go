package operations

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	logs []string
	ret  int
}

func TestOperations(t *testing.T) {
	suites := []suite{
		{
			[]string{"d1/", "d2/", "../", "d21/", "./"},
			2,
		},
		{
			[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			3,
		},
		{
			[]string{"d1/", "../", "../", "../"},
			0,
		},
	}

	for _, s := range suites {
		ret := minOperations(s.logs)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
