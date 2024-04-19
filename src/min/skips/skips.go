package skips

type f struct {
	n int
	d int
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func (num f) lt(other f) bool {
	return num.n*other.d < num.d*other.n
}

func (num f) lte(other f) bool {
	return num.n*other.d <= num.d*other.n
}

func (num f) add(other f) f {
	n := num.n*other.d + num.d*other.n
	if n == 0 {
		return f{0, 1}
	}
	d := num.d * other.d
	g := gcd(n, d)
	return f{n / g, d / g}
}

func minSkips(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	dp := make([]f, 0, n)
	dp = append(dp, f{0, 1})
	fhoursBefore := f{hoursBefore, 1}
	for _, d := range dist {
		for j := len(dp) - 1; j >= 0; j-- {
			hour := dp[j]
			temp := hour.add(f{d, speed})
			if temp.n%temp.d == 0 {
				dp[j] = temp
			} else {
				dp[j] = f{temp.n/temp.d + 1, 1}
				if j+1 < len(dp) {
					if temp.lt(dp[j+1]) {
						dp[j+1] = temp
					}
				} else {
					dp = append(dp, temp)
				}
			}
		}
	}

	for i, hours := range dp {
		if hours.lte(fhoursBefore) {
			return i
		}
	}
	return -1
}
