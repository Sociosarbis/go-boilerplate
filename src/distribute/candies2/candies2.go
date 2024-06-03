package candies2

func distributeCandies(candies int, num_people int) []int {
	num := 1
	i := 0
	ret := make([]int, num_people)
	for candies != 0 {
		if candies <= num {
			ret[i] += candies
			break
		} else {
			ret[i] += num
			candies -= num
			num++
		}
		i++
		if i == num_people {
			i = 0
		}
	}
	return ret
}
