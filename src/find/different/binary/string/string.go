package string

type node struct {
	children [2]*node
}

func find(root *node, out *[]byte, n int) bool {
	if len(*out) == n {
		return false
	}
	for i := 0; i < 2; i++ {
		*out = append(*out, byte(i)+'0')
		if root.children[i] == nil {
			for len(*out) < n {
				*out = append(*out, '0')
			}
			return true
		} else {
			if find(root.children[i], out, n) {
				return true
			}
		}
		*out = (*out)[:len(*out)-1]
	}
	return false
}

func findDifferentBinaryString(nums []string) string {
	n := len(nums[0])
	root := &node{}
	for _, num := range nums {
		cur := root
		for _, c := range num {
			index := c - '0'
			if cur.children[index] == nil {
				cur.children[index] = &node{}
			}
			cur = cur.children[index]
		}
	}
	out := make([]byte, 0, n)
	find(root, &out, n)
	return string(out)
}
