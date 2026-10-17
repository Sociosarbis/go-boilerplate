package counts2

func lowBit(index int) int {
	return index & -index
}

type bit struct {
	size int
	// 差分数组，表示[j,i]去重元素的数量增减
	tree1 []int64
	// 差分数组求实际前缀和[0,x]公式的其中一个组成部分
	// sum((x - k + 1) * bit[k]) = (x + 1) * bit[k] - k * bit[k]
	// tree2就是k * bit[k]
	tree2 []int64
}

func newBit(n int) bit {
	return bit{
		n,
		make([]int64, n+2),
		make([]int64, n+2),
	}
}

// add和query 由于分别是加/减最右的1
// 跳跃的区间分别是向右或者向左越来越大
// 所以不存在add会导致query的结果有重复添加
func (self bit) add(index, delta int) {
	index++
	rawIndex := int64(index)
	for index <= self.size {
		self.tree1[index] += int64(delta)
		self.tree2[index] += rawIndex * int64(delta)
		index += lowBit(index)
	}
}

func (self bit) addDiff(left, right, delta int) {
	self.add(left, delta)
	self.add(right+1, -delta)
}

func (self bit) queryPrefix(index int) int64 {
	index++
	rawIndex := int64(index + 1)
	if index > self.size {
		index = self.size
	}
	var res int64
	for index >= 0 {
		res += rawIndex*self.tree1[index] - self.tree2[index]
		index -= lowBit(index)
	}
	return res
}

func (self bit) query(left, right int) int64 {
	return self.queryPrefix(right) - self.queryPrefix(left-1)
}

const mod int64 = 1e9 + 7

func sumCounts(nums []int) int {
	n := len(nums)
	var res, cur int64
	b := newBit(n + 10)
	seen := map[int]int{}
	for i, num := range nums {
		var j int
		if prev, ok := seen[num]; ok {
			j = prev + 1
		}
		// j到i这段的去重元素数量会加1
		// 求j到i间的去重元素数量的和
		q := b.query(j, i)
		// 以i为end的子数组的去重元素数量的平方和
		// (xi + 1) ^ 2 - xi ^ 2 = 2 * xi - 1
		// 因为有i - j + 1项
		// 所以sum(2 * xi - 1) = 2 * q + i - j + 1
		cur = (cur + 2*q + int64(i-j) + 1) % mod
		res = (res + cur) % mod
		b.addDiff(j, i, 1)
		seen[num] = i
	}
	return int(res)
}
