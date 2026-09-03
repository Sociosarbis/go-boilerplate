package array2

const placeholder int = 1e9 + 1

func uniformArray(nums1 []int) bool {
	var a, b int = placeholder, placeholder
	for _, num := range nums1 {
		if num&1 == 1 {
			if num < b {
				b = num
			}
		} else {
			if num < a {
				a = num
			}
		}
	}
	return a == placeholder || b == placeholder || a > b
}
