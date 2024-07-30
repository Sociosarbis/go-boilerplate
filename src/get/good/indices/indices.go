package indices

func getPowerMod(a, b, c int) int {
	set := make(map[int]int)
	temp := 1
	c -= 1
	var nums []int
	for {
		temp = (temp * a) % b
		i, ok := set[temp]
		if !ok {
			set[temp] = len(nums)
			nums = append(nums, temp)
			if c < len(nums) {
				return nums[c]
			}
		} else {
			return nums[i+(c-len(nums))%(len(nums)-i)]
		}
	}
}

func getGoodIndices(variables [][]int, target int) []int {
	var ret []int
	for i, variable := range variables {
		a, b, c, m := variable[0], variable[1], variable[2], variable[3]
		if getPowerMod(getPowerMod(a, 10, b), m, c) == target {
			ret = append(ret, i)
		}
	}
	return ret
}
