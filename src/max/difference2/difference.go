package difference2

func maxDifference(s string, k int) int {
	n := len(s)
	var a, b byte
	out := -n
	for a = '0'; a <= '4'; a++ {
		for b = '0'; b <= '4'; b++ {
			if a != b {
				queues := [4][][3]int{}
				mins := [4]int{0, n, n, n}
				var aCount, bCount int
				for i := 0; i < n; i++ {
					if s[i] == a {
						aCount++
					} else if s[i] == b {
						bCount++
					}
					c := ((aCount & 1) << 1) | (bCount & 1)
					var d int
					if c < 2 {
						d = c + 2
					} else {
						d = c - 2
					}
					if aCount != 0 && bCount != 0 && i+1 >= k {
						var j int
						for len(queues[d]) > j && i+1-queues[d][j][2] >= k && queues[d][j][0] < aCount && queues[d][j][1] < bCount {
							if queues[d][j][0]-queues[d][j][1] < mins[d] {
								mins[d] = queues[d][j][0] - queues[d][j][1]
							}
							j++
						}
						queues[d] = queues[d][j:]
						if aCount-bCount-mins[d] > out {
							out = aCount - bCount - mins[d]
						}
					}
					if s[i] == a || s[i] == b {
						queues[c] = append(queues[c], [3]int{aCount, bCount, i + 1})
					}
				}
			}
		}
	}
	return out
}
