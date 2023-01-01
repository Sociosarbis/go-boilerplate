package character

func repeatedCharacter(s string) byte {
	mask := 0

	for _, c := range s {
		if mask&(1<<(c-97)) != 0 {
			return byte(c)
		}
		mask |= 1 << (c - 97)
	}
	return 0
}
