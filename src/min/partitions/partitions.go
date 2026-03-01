package partitions

func minPartitions(n string) int {
	var out rune
	for _, c := range n {
		if c > out {
			out = c
		}
	}
	return int(out - '0')
}
