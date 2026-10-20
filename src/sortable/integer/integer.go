package integer

import "sort"

func kmp(nums1 []int, nums2 []int) bool {
	m := len(nums2)
	n := m * 2
	// [0,i]的字符串里，前缀跟后缀相同的最大长度
	pi := make([]int, m)
	var j int
	for i := 1; i < m; i++ {
		for j > 0 && nums2[i] != nums2[j] {
			j = pi[j-1]
		}
		if nums2[i] == nums2[j] {
			j++
		}
		pi[i] = j
	}
	j = 0
	for i := 0; i < n; i++ {
		index := i % m
		// 如果当前的index无法匹配当前的j，则尝试逐步回溯至最大的前后缀相同的状态
		for j > 0 && nums1[index] != nums2[j] {
			j = pi[j-1]
		}
		if nums1[index] == nums2[j] {
			j++
		}
		if j == m {
			return true
		}
	}
	return false
}

func sortableIntegers(nums []int) int {
	n := len(nums)
	target := make([]int, n)
	copy(target, nums)
	sort.Ints(target)
	var out int
	for x := 1; x*x <= n; x++ {
		if n%x != 0 {
			continue
		}
		d := x
	block:
		for true {
			for left := 0; left < n; left += d {
				if !kmp(nums[left:left+d], target[left:left+d]) {
					break block
				}
			}
			out += d
			break
		}
		d = n / x
		if d != x {
		block2:
			for true {
				for left := 0; left < n; left += d {
					if !kmp(nums[left:left+d], target[left:left+d]) {
						break block2
					}
				}
				out += d
				break
			}
		}
	}
	return out
}
