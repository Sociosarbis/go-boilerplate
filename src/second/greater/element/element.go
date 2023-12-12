package element

type item struct {
	i int
	v int
}

func secondGreaterElement(nums []int) []int {
	ret := make([]int, len(nums))
	for i := range ret {
		ret[i] = -1
	}
	level2 := []item{}
	level1 := []item{}

	for i, num := range nums {
		j := len(level2)
		for ; j > 0; j-- {
			if level2[j-1].v < num {
				ret[level2[j-1].i] = num
			} else {
				break
			}
		}
		level2 = level2[:j]
		j = len(level1)
		for ; j > 0; j-- {
			if level1[j-1].v >= num {
				break
			}
		}
		level2 = append(level2, level1[j:]...)
		level1 = level1[:j]
		level1 = append(level1, item{
			i,
			num,
		})
	}
	return ret
}
