package number

func nextBeautifulNumber(n int) int {
	if n == 0 {
		return 1
	}
	target := make([]int, 0, 7)
	stocks := make([]int, 7)
	for i := 1; i < len(stocks); i++ {
		stocks[i] = i
	}
	for n != 0 {
		target = append(target, n%10)
		n /= 10
	}
	res := dfs(target, len(target)-1, make([]int, 0, 7), false, stocks, 0)
	ret := 0
	base := 1
	for i := len(res) - 1; i >= 0; i-- {
		ret += res[i] * base
		base *= 10
	}
	return ret
}

func dfs(target []int, i int, temp []int, isGreater bool, stocks []int, selected int) []int {
	if i < 0 {
		if isGreater {
			return temp
		} else {
			return nil
		}
	}
	targetDigit := target[i]
	for j := 1; j < len(stocks); j++ {
		if stocks[j] > 0 && (isGreater || j >= targetDigit) && (stocks[j] != j || selected+j <= len(target)) {
			temp = append(temp, j)
			nextSelected := selected
			if stocks[j] == j {
				nextSelected += j
			}
			stocks[j] -= 1
			nextGreater := isGreater
			if !nextGreater && j > targetDigit {
				nextGreater = true
			}
			res := dfs(target, i-1, temp, nextGreater, stocks, nextSelected)
			if res != nil {
				return res
			}
			temp = temp[:len(temp)-1]
			stocks[j] += 1
		}
	}
	if i+1 == len(target) {
		target = make([]int, len(target)+1)
		target[len(target)-1] = 1
		return dfs(target, i+1, temp, false, stocks, selected)
	}
	return nil
}
