package possibility

import (
	"testing"
)

type suite struct {
	nums []int
	ret  bool
}

func TestPossibility(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 2, 3},
			true,
		},
		{
			[]int{4, 2, 1},
			false,
		},
	}

	for _, s := range suites {
		ret := checkPossibility(s.nums)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
