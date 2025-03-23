package valid

func canBeValid(s string, locked string) bool {
	var count int
	var leftCount int
	var leftClosed int
	n := len(s)
	for i := 0; i < n; i++ {
		if locked[i] == '1' {
			if s[i] == ')' {
				if leftCount != 0 {
					leftCount--
				} else {
					if count > 0 {
						count--
					} else if leftClosed > 0 {
						leftClosed--
						leftCount++
					} else {
						return false
					}
				}
			} else {
				leftCount++
			}
		} else {
			if leftCount > 0 {
				leftCount--
				leftClosed++
			} else {
				count++
			}
		}
	}
	return leftCount == 0 && count&1 == 0
}
