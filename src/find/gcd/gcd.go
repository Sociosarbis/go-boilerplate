package gcd

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func findGCD(nums []int) int {
	a := 1001
	var b int
	for _, num := range nums {
		if num < a {
			a = num
		}
		if num > b {
			b = num
		}
	}
	return gcd(a, b)
}
