package secret

func normalize(num int) int {
	if num < 0 {
		return num + 1e9 + 7
	}
	return num % (1e9 + 7)
}

func peopleAwareOfSecret(n int, delay int, forget int) int {
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp2[delay] = 1
	if forget < n {
		dp1[forget] = 1
	}
	var canShare int
	out := 1
	for i := 1; i < n; i++ {
		canShare = normalize(canShare - dp1[i])
		out = normalize(out - dp1[i])
		canShare = normalize(canShare + dp2[i])
		out = normalize(out + canShare)
		if i+delay < n {
			dp2[i+delay] = canShare
		}
		if i+forget < n {
			dp1[i+forget] = canShare
		}
	}
	return out
}
