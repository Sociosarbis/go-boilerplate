package lhs

func findLHS(nums []int) int {
	counter := map[int]int{}
	for _, num := range nums {
		if _, ok := counter[num]; !ok {
			counter[num] = 0
		}
		counter[num] += 1
	}
	ret := 0
	for num := range counter {
		if count2, ok := counter[num+1]; ok {
			count1 := counter[num]
			if count1+count2 > ret {
				ret = count1 + count2
			}
		}
	}
	return ret
}
