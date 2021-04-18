package duplicates

import (
	"testing"
)

type Suite struct {
	nums []int
	ret  int
}

func TestRemoveDuplicates(t *testing.T) {
	suites := []Suite{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}

	for _, s := range suites {
		ret := removeDuplicates(s.nums)
		if ret != s.ret {
			t.Errorf("Failed, expected %v but get %v", s.ret, ret)
		}
	}
}
