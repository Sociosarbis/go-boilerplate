package length5

const maxPossibleNumGreaterThan1 int = 31622

func maximumLength(nums []int) int {
	n := len(nums)
	m := make(map[int]byte, n)
	var ones int
	for _, num := range nums {
		if num == 1 {
			ones++
			continue
		}
		if c, ok := m[num]; ok {
			if c < 2 {
				m[num] = c + 1
			}
		} else {
			m[num] = 1
		}
	}
	var out int
	if ones > 0 {
		if ones&1 == 1 {
			out = ones
		} else {
			out = ones - 1
		}
	} else {
		out = 1
	}
	var temp int
	for num, count := range m {
		if count < 2 {
			continue
		}
		temp = 0
		for num <= maxPossibleNumGreaterThan1 {
			temp++
			num *= num
			if c, ok := m[num]; ok {
				if temp*2+1 > out {
					out = temp*2 + 1
				}
				if c < 2 {
					break
				}
			} else {
				if temp*2-1 > out {
					out = temp*2 - 1
				}
				break
			}
		}
	}
	return out
}
