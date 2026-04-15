package target

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func distance(i, j, n int) int {
	var out int
	if i < j {
		out = min(j-i, i+n-j)
	} else {
		out = min(i-j, j+n-i)
	}
	return out
}

func closestTarget(words []string, target string, startIndex int) int {
	n := len(words)
	out := n
	for i := 0; i < n; i++ {
		if words[i] == target {
			temp := distance(i, startIndex, n)
			if temp < out {
				out = temp
			}
		}
	}
	if out == n {
		return -1
	}
	return out
}
