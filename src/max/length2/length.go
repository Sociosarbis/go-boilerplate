package length2

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func maxLength(nums []int) int {
	var l int
	n := len(nums)
	out := 1
	p, la, ga := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		nla, nga := lcm(la, nums[i]), gcd(ga, nums[i])
		np := p * nums[i]
		if nla*nga != np {
			if i-l > out {
				out = i - l
			}
			l = i
			p, la, ga = nums[i], nums[i], nums[i]
			for l > 0 {
				nla, nga := lcm(la, nums[l-1]), gcd(ga, nums[l-1])
				np := p * nums[l-1]
				if nla*nga != np {
					break
				}
				l--
				la, ga = nla, nga
				p = np
			}
		} else {
			la, ga = nla, nga
			p = np
		}
	}
	if n-l > out {
		return n - l
	}
	return out
}
