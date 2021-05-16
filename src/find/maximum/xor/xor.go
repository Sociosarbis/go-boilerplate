package xor

type tree struct {
	left  *tree
	right *tree
}

func findMaximumXOR(nums []int) int {
	root := tree{}
	max_bit := 1 << 30
	ret := 0
	for _, num := range nums {
		cur_bit := max_bit
		cur_node := &root
		for cur_bit != 0 {
			if cur_bit&num != 0 {
				if cur_node.right == nil {
					cur_node.right = &tree{}
				}
				cur_node = cur_node.right
			} else {
				if cur_node.left == nil {
					cur_node.left = &tree{}
				}
				cur_node = cur_node.left
			}
			cur_bit >>= 1
		}

		cur_bit = max_bit
		temp := 0
		cur_node = &root
		for cur_bit != 0 {
			if cur_bit&num != 0 {
				if cur_node.left != nil {
					temp += cur_bit
					cur_node = cur_node.left
				} else {
					cur_node = cur_node.right
				}
			} else {
				if cur_node.right != nil {
					temp += cur_bit
					cur_node = cur_node.right
				} else {
					cur_node = cur_node.left
				}
			}
			cur_bit >>= 1
		}
		if temp > ret {
			ret = temp
		}
	}
	return ret
}
