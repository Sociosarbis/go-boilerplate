package water

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	mem := map[int]map[int]struct{}{}
	return dfs(jug1Capacity, jug2Capacity, 0, 0, targetCapacity, mem)
}

func dfs(jug1Capacity, jug2Capacity, temp1, temp2, targetCapacity int, mem map[int]map[int]struct{}) bool {
	if temp1+temp2 == targetCapacity {
		return true
	}
	if _, ok := mem[temp1]; !ok {
		mem[temp1] = map[int]struct{}{}
	}

	if _, ok := mem[temp1][temp2]; ok {
		return false
	} else {
		mem[temp1][temp2] = struct{}{}
	}
	if dfs(jug1Capacity, jug2Capacity, jug1Capacity, temp2, targetCapacity, mem) {
		return true
	}
	if dfs(jug1Capacity, jug2Capacity, temp1, jug2Capacity, targetCapacity, mem) {
		return true
	}
	if dfs(jug1Capacity, jug2Capacity, 0, temp2, targetCapacity, mem) {
		return true
	}
	if dfs(jug1Capacity, jug2Capacity, temp1, 0, targetCapacity, mem) {
		return true
	}
	sum := temp1 + temp2
	if sum <= jug1Capacity {
		if dfs(jug1Capacity, jug2Capacity, sum, 0, targetCapacity, mem) {
			return true
		}
	} else {
		if dfs(jug1Capacity, jug2Capacity, jug1Capacity, sum-jug1Capacity, targetCapacity, mem) {
			return true
		}
	}
	if sum <= jug2Capacity {
		if dfs(jug1Capacity, jug2Capacity, 0, sum, targetCapacity, mem) {
			return true
		}
	} else {
		if dfs(jug1Capacity, jug2Capacity, sum-jug2Capacity, jug2Capacity, targetCapacity, mem) {
			return true
		}
	}
	return false
}
