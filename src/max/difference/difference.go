package difference

func maxDifference(s string) int {
	counter := [26]int{}
	for _, c := range s {
		counter[c-97]++
	}
	var a int
	b := 100
	for i := 0; i < 26; i++ {
		if counter[i] != 0 {
			if counter[i]&1 == 1 {
				if counter[i] > a {
					a = counter[i]
				}
			} else {
				if counter[i] < b {
					b = counter[i]
				}
			}
		}
	}
	return a - b
}
