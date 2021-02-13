package numbers

import (
	"reflect"
	"testing"
)

type suite struct {
	nums []int
	ret  []int
}

func TestNumbers(t *testing.T) {
	suites := []suite{
		{
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			[]int{2, 1},
			[]int{},
		},
	}
	for _, su := range suites {
		ret := findDisappearedNumbers(su.nums)
		if !reflect.DeepEqual(ret, su.ret) {
			t.Errorf("Failed, expected %v but get %v", su.ret, ret)
		}
	}
}
