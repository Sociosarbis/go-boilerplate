package operations18

func minOperations(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}
	min, max := s[0], s[0]
	mnc, mxc := 1, 1
	sorted := s[1] >= s[0]
	if s[1] > max {
		max = s[1]
	} else if s[1] == max {
		mxc++
	}
	if s[1] < min {
		min = s[1]
	} else if s[1] == min {
		mnc++
	}
	for i := 2; i < n; i++ {
		if sorted && s[i] < s[i-1] {
			sorted = false
		}
		if s[i] > max {
			max, mxc = s[i], 1
		} else if mxc < 2 && s[i] == max {
			mxc++
		}
		if s[i] < min {
			min, mnc = s[i], 1
		} else if mnc < 2 && s[i] == min {
			mnc++
		}
	}
	if sorted {
		return 0
	} else if n == 2 {
		return -1
	}
	var out int
	if s[0] == min || s[n-1] == max {
		out = 1
	} else if mnc == 1 && mxc == 1 && s[n-1] == min && s[0] == max {
		out = 3
	} else {
		out = 2
	}
	return out
}
