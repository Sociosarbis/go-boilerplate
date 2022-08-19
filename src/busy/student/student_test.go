package student

import (
	"testing"

	"github.com/boilerplate/src/constant"
)

type suite struct {
	startTime []int
	endTime   []int
	queryTime int
	ret       int
}

func TestStudent(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 7},
			4,
			1,
		},
		{
			[]int{4},
			[]int{4},
			4,
			1,
		},
	}

	for _, s := range suites {
		ret := busyStudent(s.startTime, s.endTime, s.queryTime)
		if ret != s.ret {
			t.Errorf(constant.ErrTemplateStr, s.ret, ret)
		}
	}
}
