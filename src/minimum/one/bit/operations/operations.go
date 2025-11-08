package operations

func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}
	if n&(n-1) == 0 {
		return n + (n - 1)
	}
	var i int
	temp := n
	for temp != 0 {
		temp >>= 1
		i++
	}
	mse := 1 << (i - 1)
	// i是n的最高置位，1-base
	// 从2^(i - 1)转到0都会经过最高位为i的任意n，所以可以做如下递归
	return minimumOneBitOperations(mse) - minimumOneBitOperations(n-mse)
}
