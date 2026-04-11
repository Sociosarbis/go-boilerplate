package distance3

type group struct {
	nums   [3]int
	length int
}

func (g *group) add(num int) {
	if g.length < 3 {
		g.nums[g.length] = num
		g.length++
	} else {
		g.nums[0], g.nums[1], g.nums[2] = g.nums[1], g.nums[2], num
	}
}

func (g *group) dist() int {
	if g.length == 3 {
		return g.nums[2] - g.nums[0]
	}
	return 1e5
}

func minimumDistance(nums []int) int {
	numToIndices := make([]group, len(nums)+1)
	var out int = 1e5
	for i, num := range nums {
		numToIndices[num].add(i)
		if numToIndices[num].dist() < out {
			out = numToIndices[num].dist()
		}
	}
	if out == 1e5 {
		return -1
	}
	return out * 2
}
