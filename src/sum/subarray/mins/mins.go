package mins

type item struct {
	i int
	v int
}

func sumSubarrayMins(arr []int) int {
	stack := []item{
		{
			0,
			arr[0],
		},
	}
	arr = append(arr, 0)
	ret := 0
	var mask int = 1e9 + 7
	for i := 1; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] < stack[len(stack)-1].v {
			temp := i - stack[len(stack)-1].i
			if len(stack) == 1 {
				temp *= stack[len(stack)-1].i + 1
			} else {
				temp *= stack[len(stack)-1].i - stack[len(stack)-2].i
			}
			ret = (ret + temp*stack[len(stack)-1].v) % mask
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, item{i, arr[i]})
	}
	return ret
}
