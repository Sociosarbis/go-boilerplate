package area

import "sort"

func maxArea(height []int) int {
	n := len(height)
	stack := make([][2]int, 0, n)
	stack = append(stack, [2]int{0, height[0]})
	var out int
	for i := 1; i < n; i++ {
		index := sort.Search(len(stack), func(j int) bool {
			return stack[j][1] >= height[i]
		})
		if index != len(stack) {
			temp := (i - stack[index][0]) * height[i]
			if temp > out {
				out = temp
			}
		}
		if stack[len(stack)-1][1] < height[i] {
			stack = append(stack, [2]int{i, height[i]})
		}
	}
	stack = append(stack[:0], [2]int{n - 1, height[n-1]})
	for i := n - 2; i >= 0; i-- {
		index := sort.Search(len(stack), func(j int) bool {
			return stack[j][1] >= height[i]
		})
		if index != len(stack) {
			temp := (stack[index][0] - i) * height[i]
			if temp > out {
				out = temp
			}
		}
		if stack[len(stack)-1][1] < height[i] {
			stack = append(stack, [2]int{i, height[i]})
		}
	}
	return out
}
