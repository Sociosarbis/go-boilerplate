package reach2

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] != '0' {
		return false
	}
	near := []int{0}
	from := []int{}
	for i := 1; i < n; i++ {
		var j int
		for ; j < len(near); j++ {
			if i-near[j] < minJump {
				break
			}
		}
		from = append(from, near[:j]...)
		near = near[j:]
		j = 0
		for ; j < len(from); j++ {
			if i-from[j] <= maxJump {
				break
			}
		}
		from = from[j:]
		if len(near) == 0 && len(from) == 0 {
			return false
		}
		if s[i] == '0' && len(from) > 0 {
			near = append(near, i)
		}
	}
	return len(from) > 0
}
