package weeks

func numberOfWeeks(milestones []int) int64 {
	n := len(milestones)
	max := milestones[0]
	sum := int64(max)
	for i := 1; i < n; i++ {
		if milestones[i] > max {
			max = milestones[i]
		}
		sum += int64(milestones[i])
	}
	rest := sum - int64(max)
	// 当rest至多比max少1时，max和非max两两配对肯定能消耗完max
	// 当max消耗完以后，按最优的方式消耗的话，下个max肯定也会符合rest至多比max少1的条件
	// 这样递归到只剩2个时，此时要符合rest至多比max少1的条件
	// 非max的另一个数必然大于等于max - 1，此时必然都能被消耗
	if rest >= int64(max-1) {
		return sum
	}
	return 2*rest + 1
}
