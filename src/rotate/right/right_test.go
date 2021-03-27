package right

import (
	"reflect"
	"testing"
)

type Suite struct {
	head []int
	k    int
	ret  []int
}

func TestRight(t *testing.T) {
	suites := []Suite{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{4, 5, 1, 2, 3},
		},
		{
			[]int{0, 1, 2},
			4,
			[]int{2, 0, 1},
		},
		{
			[]int{1, 2},
			1,
			[]int{2, 1},
		},
	}

	for i := 0; i < len(suites); i++ {
		input := ListNode{
			0,
			nil,
		}
		cur := &input
		for j := 0; j < len(suites[i].head); j++ {
			cur.Next = &ListNode{
				suites[i].head[j],
				nil,
			}
			cur = cur.Next
		}
		ret := rotateRight(input.Next, suites[i].k)
		output := []int{}
		for ret != nil {
			output = append(output, ret.Val)
			ret = ret.Next
		}

		if !reflect.DeepEqual(output, suites[i].ret) {
			t.Errorf("Failed, expected %v but get %v", suites[i].ret, output)
		}
	}
}
