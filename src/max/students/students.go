package students

func maxStudents(seats [][]byte) int {
	n := len(seats[0])
	dp := make([]int, 1<<n)
	ret := 0
	prevBlock := 0
	for _, row := range seats {
		nextDp := make([]int, 1<<n)
		nextDp[0] = ret
		block := 0
		for j, cell := range row {
			if cell != '#' {
				end := 1
				if j > 0 {
					end = 1 << (j - 1)
				}
				for index := 0; index < end; index++ {
					if index&(block|(index<<1)|(index>>1)) == 0 {
						idx := index | (1 << j)
						not_mask := (idx << 1) | (idx >> 1) | prevBlock
						count := countOne(index)
						for k, num := range dp {
							if k&not_mask == 0 {
								temp := num + count + 1
								if temp > nextDp[idx] {
									nextDp[idx] = temp
									if nextDp[idx] > ret {
										ret = nextDp[idx]
									}
								}
							}
						}
					}
				}
			} else {
				block |= 1 << j
			}
		}
		dp = nextDp
		prevBlock = block
	}

	return ret
}

func countOne(num int) int {
	ret := 0
	for num > 0 {
		num &= num - 1
		ret++
	}
	return ret
}
