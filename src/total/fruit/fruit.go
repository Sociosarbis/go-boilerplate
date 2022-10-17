package fruit

func totalFruit(fruits []int) int {
	types := make(map[int]bool, 2)
	ret := 0
	temp := 0
	lastChunk := make([]int, 2)
	for i := 0; i < len(fruits); i++ {
		chunk := make([]int, 2)
		chunk[0] = fruits[i]
		j := i + 1
		for ; j < len(fruits); j++ {
			if fruits[j] != fruits[i] {
				break
			}
		}
		chunk[1] = j - i
		if len(types) < 2 {
			if _, ok := types[fruits[i]]; !ok {
				types[fruits[i]] = true
			}
		} else {
			if _, ok := types[fruits[i]]; !ok {
				if temp > ret {
					ret = temp
				}
				for k := range types {
					if k != lastChunk[0] {
						delete(types, k)
						break
					}
				}
				types[fruits[i]] = true
				temp = lastChunk[1]
			}
		}
		lastChunk = chunk
		temp += chunk[1]
		i = j - 1
	}
	if temp > ret {
		ret = temp
	}
	return ret
}
