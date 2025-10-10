package energy

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	for i := n - k - 1; i >= 0; i-- {
		energy[i] += energy[i+k]
	}
	out := energy[0]
	for i := 1; i < n; i++ {
		if energy[i] > out {
			out = energy[i]
		}
	}
	return out
}
