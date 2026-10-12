package sum2

const mask int64 = 1e9 + 7

func maxSum(nums []int, k int) int {
	var counter [31]int
	for _, num := range nums {
		var i int
		for num != 0 {
			if num&1 == 1 {
				counter[i]++
			}
			num >>= 1
			i++
		}
	}
	n := len(nums)
	var out int64
	for k > 0 {
		min := n
		var value int
		for i := 0; i < 31; i++ {
			if counter[i] != 0 {
				if counter[i] < min {
					min = counter[i]
				}
				value |= 1 << i
			}
		}
		if min > k {
			min = k
		} else {
			for i := 0; i < 31; i++ {
				if counter[i] != 0 {
					counter[i] -= min
				}
			}
		}
		out = (out + (((int64(value)*int64(value))%mask)*int64(min))%mask) % mask
		k -= min
	}
	return int(out)
}
