package win2

func canAliceWin(n int) bool {
	for i := 10; i >= 1; i-- {
		if n < i {
			if i&1 == 1 {
				return true
			} else {
				return false
			}
		}
		n -= i
	}
	return false
}
