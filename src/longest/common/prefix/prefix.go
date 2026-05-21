package prefix

type node struct {
	next []*node
}

func add(num int, cur *node) {
	digits := make([]int, 0, 9)
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if cur.next == nil {
			cur.next = make([]*node, 10)
		}
		if cur.next[digits[i]] == nil {
			cur.next[digits[i]] = &node{}
		}
		cur = cur.next[digits[i]]
	}
}

func count(num int, cur *node) int {
	digits := make([]int, 0, 9)
	for num != 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	n := len(digits)
	var out int
	for i := n - 1; i >= 0; i-- {
		if cur.next != nil && cur.next[digits[i]] != nil {
			out++
			cur = cur.next[digits[i]]
		} else {
			break
		}
	}
	return out
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	var out int
	root := &node{}
	for _, num := range arr1 {
		add(num, root)
	}
	for _, num := range arr2 {
		res := count(num, root)
		if res > out {
			out = res
		}
	}
	return out
}
