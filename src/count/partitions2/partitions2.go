package partitions2

import "sort"

const mask int = 1e9 + 7

func minus(a, b int) int {
	if a < b {
		return a - b + mask
	}
	return a - b
}

func countPartitions(nums []int, k int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	minStack := [][2]int{{nums[0], 0}}
	maxStack := [][2]int{{nums[0], 0}}
	prefixSum[1] = 1
	start := -1
	for i := 1; i < n; i++ {
		num := nums[i]
		j := len(minStack) - 1
		index := sort.Search(len(minStack), func(i int) bool {
			return minStack[i][0] >= num-k
		})
		if index > 0 {
			if minStack[index-1][1] > start {
				start = minStack[index-1][1]
			}
		}
		index = sort.Search(len(maxStack), func(i int) bool {
			return maxStack[i][0] <= num+k
		})
		if index > 0 {
			if maxStack[index-1][1] > start {
				start = maxStack[index-1][1]
			}
		}
		if start == -1 {
			prefixSum[i+1] = 1
			index = 0
		} else {
			index = start
		}
		prefixSum[i+1] = (prefixSum[i+1] + prefixSum[i] + minus(prefixSum[i], prefixSum[index])) % mask
		for ; j >= 0; j-- {
			if minStack[j][0] < num {
				break
			}
		}
		minStack = append(minStack[:j+1], [2]int{num, i})
		j = len(maxStack) - 1
		for ; j >= 0; j-- {
			if maxStack[j][0] > num {
				break
			}
		}
		maxStack = append(maxStack[:j+1], [2]int{num, i})
	}
	return minus(prefixSum[n], prefixSum[n-1])
}
