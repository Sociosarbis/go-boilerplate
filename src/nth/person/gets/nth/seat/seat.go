package seat

// 对于n > 2
// f(n) = 1/n + 0 + (n-2)/n * (1/(n-2) * f(n-1) + 1/(n-2) * f(n-2) + ... + 1/(n-2) * f(2))
// f(n) = 1/n + 1/n * (f(n-1) + f(n-2) + ... + f(2))
// n * f(n) = 1 + f(n-1) + f(n-2) + ... + f(2)
// (n-1) * f(n-1) = 1 + f(n-2) + f(n-3) + ... + f(2)
// n * f(n) - (n-1) * f(n-1) = f(n-1)
// f(n) = f(n-1)
// 关键是理解对于每个位置，假如剩下n个位置，他都有三种情况，选到一个不属于后面的位置，选到最后一个位置
// 选到除上述两个位置的位置，对于这种情况，每个位置k选不到最后一个位置的概率都是1 / (n - 2) * f(n - k)

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}