package swaps3

import "sort"

func minSwaps(nums []int) int {
	n := len(nums)
	digitSums := make([][2]int, n)
	for i, num := range nums {
		var temp int
		for num != 0 {
			temp += num % 10
			num /= 10
		}
		digitSums[i][0], digitSums[i][1] = i, temp
	}
	sort.Slice(digitSums, func(i, j int) bool {
		if digitSums[i][1] < digitSums[j][1] {
			return true
		}
		if digitSums[i][1] > digitSums[j][1] {
			return false
		}
		return nums[digitSums[i][0]] < nums[digitSums[j][0]]
	})
	oldToNew := make([]int, n)
	for i, digitSum := range digitSums {
		oldToNew[digitSum[0]] = i
	}
	var cycles int
	visited := make([]bool, n)
	for next, digitSum := range digitSums {
		if visited[next] {
			continue
		}
		start := digitSum[0]
		visited[next] = true
		for next != start {
			next = oldToNew[next]
			visited[next] = true
		}
		cycles++
	}
	// 最小交换次数为n - 环数量
	// 因为每个环变成正确的顺序，需要环成员数-1次交换
	return n - cycles
}
