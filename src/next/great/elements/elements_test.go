package elements

import (
	"reflect"
	"testing"
)

type Suite struct {
	nums []int
	ret  []int
}

func TestNextGreatElements(t *testing.T) {
	suites := []Suite{
		{[]int{1, 2, 1},
			[]int{2, -1, 2}},
	}

	for i := 0; i < len(suites); i++ {
		ret := nextGreaterElements(suites[i].nums)

		if !reflect.DeepEqual(ret, suites[i].ret) {
			t.Errorf("Failed, expected %v but get %v", suites[i].ret, ret)
		}
	}
}
