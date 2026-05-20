package array

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	count := make([]byte, n+1)
	out := make([]int, n)
	var temp int
	for i := 0; i < n; i++ {
		count[A[i]]++
		count[B[i]]++
		if count[A[i]] == 2 {
			temp++
		}
		if A[i] != B[i] && count[B[i]] == 2 {
			temp++
		}
		out[i] = temp
	}
	return out
}
