package codes

func hasAllCodes(s string, k int) bool {
	count := 1 << k
	mask := count - 1
	n := len(s)
	if n-k+1 < count {
		return false
	}
	var temp int
	if s[0] == '1' {
		temp = 1
	}
	for i := 1; i < k; i++ {
		temp <<= 1
		if s[i] == '1' {
			temp++
		}
	}
	visited := make([]bool, count)
	visited[temp] = true
	count--
	for i := k; i < n && count > 0; i++ {
		temp = (temp << 1) & mask
		if s[i] == '1' {
			temp += 1
		}
		if !visited[temp] {
			count--
			visited[temp] = true
		}
	}
	return count == 0
}
