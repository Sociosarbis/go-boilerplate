package students

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	students   []int
	sandwiches []int
	ret        int
}

func TestStudents(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 1, 0, 0},
			[]int{0, 1, 0, 1},
			0,
		},
		{
			[]int{1, 1, 1, 0, 0, 1},
			[]int{1, 0, 0, 0, 1, 1},
			3,
		},
	}

	for _, s := range suites {
		ret := countStudents(s.students, s.sandwiches)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
