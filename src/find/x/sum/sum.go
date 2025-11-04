package sum

import "sort"

func findXSum(nums []int, k int, x int) []int {
	queue := make([]int, 0, k)
	visited := make([]bool, 51)
	counter := make([]int, 51)
	for i := 0; i < k; i++ {
		counter[nums[i]]++
		if !visited[nums[i]] {
			visited[nums[i]] = true
			queue = append(queue, nums[i])
		}
	}
	sort.Slice(queue, func(i, j int) bool {
		return counter[queue[i]] > counter[queue[j]] || (counter[queue[i]] == counter[queue[j]] && queue[i] > queue[j])
	})
	n := len(nums)
	out := make([]int, n-k+1)
	n = len(out)
	var l int
	if len(queue) > x {
		l = x
	} else {
		l = len(queue)
	}
	for i := 0; i < l; i++ {
		out[0] += counter[queue[i]] * queue[i]
	}
	for i := 1; i < n; i++ {
		counter[nums[i-1]]--
		counter[nums[i+k-1]]++
		if !visited[nums[i+k-1]] {
			visited[nums[i+k-1]] = true
			queue = append(queue, nums[i+k-1])
		}
		sort.Slice(queue, func(i, j int) bool {
			return counter[queue[i]] > counter[queue[j]] || (counter[queue[i]] == counter[queue[j]] && queue[i] > queue[j])
		})
		var l int
		if len(queue) > x {
			l = x
		} else {
			l = len(queue)
		}
		for j := 0; j < l; j++ {
			out[i] += counter[queue[j]] * queue[j]
		}
	}
	return out
}
