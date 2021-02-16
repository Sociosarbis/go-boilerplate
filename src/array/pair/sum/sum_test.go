package sum

import (
	"testing"
)

type suite struct {
	nums []int
	ret  int
}

func TestSum(t *testing.T) {
	suites := []suite{
		{
			[]int{1, 4, 3, 2},
			4,
		},
		{
			[]int{6, 2, 6, 5, 1, 2},
			9,
		},
	}

	for _, su := range suites {
		ret := arrayPairSum(su.nums)
		if ret != su.ret {
			t.Errorf("Failed, expected %v but get %v", su.ret, ret)
		}
	}
}
