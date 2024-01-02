package repetitions

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// 记录循环后遍历s2的累加次数
	repeatCount := make([]int, n1+1)
	// 记录下次循环开始的s2 index
	nextIndexToIter := make([]int, len(s2))
	for i := 1; i < len(s2); i++ {
		nextIndexToIter[i] = len(s2)
	}
	j := 0
	count := 0
	for i := 1; i <= n1; i++ {
		for k := 0; k < len(s1); k++ {
			if s1[k] == s2[j] {
				j++
				if j == len(s2) {
					count++
					j = 0
				}
			}
		}
		repeatCount[i] = count
		if nextIndexToIter[j] < len(s2) {
			start := nextIndexToIter[j]
			interval := i - start
			r1 := (n1 - start) / interval * (repeatCount[i] - repeatCount[start])
			// start及其之前的加上重复pattern剩余的段数
			r2 := repeatCount[start+(n1-start)%interval]
			return (r1 + r2) / n2
		}
		nextIndexToIter[j] = i
	}
	return repeatCount[n1] / n2
}
