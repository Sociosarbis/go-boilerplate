package subarrays

func countCompleteSubarrays(nums []int) int {
	counter := [2001]int{}
	visited := [2001]bool{}
	var count int
	for _, num := range nums {
		if !visited[num] {
			visited[num] = true
			count++
		}
	}
	var j int
	n := len(nums)
	var temp int
	var ret int
	for i := 0; i < n; i++ {
		if temp >= count {
			ret += n - j + 1
		} else {
			for ; j < n; j++ {
				if counter[nums[j]] == 0 {
					temp++
				}
				counter[nums[j]]++
				if temp >= count {
					j++
					ret += n - j + 1
					break
				}
			}
		}
		counter[nums[i]]--
		if counter[nums[i]] == 0 {
			temp--
		}
	}
	return ret
}
