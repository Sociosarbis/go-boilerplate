package node

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	if node1 == node2 {
		return node1
	}
	n := len(edges)
	v1, v2 := make([]bool, n), make([]bool, n)
	ret := n
	for node1 != -1 || node2 != -1 {
		if node1 == node2 {
			return node1
		}
		if node1 != -1 {
			if v1[node1] {
				node1 = -1
			} else {
				if v2[node1] {
					if node1 < ret {
						ret = node1
					}
				}
				v1[node1] = true
				node1 = edges[node1]
			}
		}
		if node2 != -1 {
			if v2[node2] {
				node2 = -1
			} else {
				if v1[node2] {
					if node2 < ret {
						ret = node2
					}
				}
				v2[node2] = true
				node2 = edges[node2]
			}
		}
		if ret != n {
			return ret
		}
	}
	return -1
}
