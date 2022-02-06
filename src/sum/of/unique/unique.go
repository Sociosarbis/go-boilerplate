package unique

func sumOfUnique(nums []int) int {
	numToIsUnique := map[int]bool{}
	for _, num := range nums {
		isUnique, ok := numToIsUnique[num]
		if !ok {
			numToIsUnique[num] = true
		} else if isUnique {
			numToIsUnique[num] = false
		}
	}

	ret := 0
	for num, isUnique := range numToIsUnique {
		if isUnique {
			ret += num
		}
	}
	return ret
}
