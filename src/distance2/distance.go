package distance2

type item struct {
	value int64
	count int
}

func distance(nums []int) []int64 {
	n := len(nums)
	counter := make(map[int]item, n)
	prefix := make([]item, n)
	for i, num := range nums {
		if it, ok := counter[num]; ok {
			it.count++
			it.value += int64(i)
			counter[num] = it
			prefix[i] = it
		} else {
			prefix[i] = item{
				count: 1,
				value: int64(i),
			}
			counter[num] = prefix[i]
		}
	}
	out := make([]int64, n)
	for i := 0; i < n; i++ {
		if it, ok := counter[nums[i]]; ok {
			out[i] = int64(i)*int64(prefix[i].count) - prefix[i].value + it.value - prefix[i].value - int64(it.count-prefix[i].count)*int64(i)
		}
	}
	return out
}
