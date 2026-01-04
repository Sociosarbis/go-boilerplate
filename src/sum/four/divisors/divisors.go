package divisors

func count(num int) int {
	if num < 6 {
		return 0
	}
	var count int
	sum := 1 + num
	for i := 2; i < num; i++ {
		prod := i * i
		if prod > num {
			break
		} else if prod == num {
			return 0
		}
		if num%i == 0 {
			count += 2
			if count > 2 {
				return 0
			}
			sum += i + num/i
		}
	}
	if count == 2 {
		return sum
	}
	return 0
}

func sumFourDivisors(nums []int) int {
	var out int
	for _, num := range nums {
		out += count(num)
	}
	return out
}
