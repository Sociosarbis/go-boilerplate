package winner

type node struct {
	v int
	n *node
}

func getWinner(arr []int, k int) int {
	n := len(arr)
	if k >= n {
		var ret int
		for _, num := range arr {
			if num > ret {
				ret = num
			}
		}
		return ret
	}
	var temp int
	cur := &node{
		arr[0],
		nil,
	}
	head := cur
	for i := 1; i < n; i++ {
		cur.n = &node{
			arr[i],
			nil,
		}
		cur = cur.n
	}
	tail := cur
	for {
		a, b := head, head.n
		if a.v > b.v {
			if b.n != nil {
				tail.n = b
				head.n = b.n
				b.n = nil
			}
			temp++
		} else {
			tail.n = a
			head = b
			a.n = nil
			temp = 1
		}
		if temp == k {
			return head.v
		}
	}
}
