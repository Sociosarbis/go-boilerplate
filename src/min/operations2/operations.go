package operation2

func minOperations(boxes string) []int {
	n := len(boxes)
	p1 := make([]int, n)
	p2 := make([]int, n)
	p3 := make([]int, n)
	if boxes[0] == '1' {
		p1[0] = 1
	}
	for i := 1; i < n; i++ {
		if boxes[i] == '1' {
			p1[i] = p1[i-1] + 1
			p2[i] = p2[i-1] + i
		} else {
			p1[i] = p1[i-1]
			p2[i] = p2[i-1]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if boxes[i] == '1' {
			p3[i] = p3[i+1] + n - 1 - i
		} else {
			p3[i] = p3[i+1]
		}
	}
	ret := make([]int, n)
	ret[0] = p2[n-1]
	for i := 1; i < n; i++ {
		ret[i] = (p3[0] - p3[i]) - (n-1-i)*p1[i-1] + (p2[n-1] - p2[i]) - (p1[n-1]-p1[i])*i
	}
	return ret
}
