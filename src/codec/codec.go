package codec

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func dfsSer(root *TreeNode, ret *string) {
	if root != nil {
		if len(*ret) != 0 {
			*ret += ","
		}
		*ret += fmt.Sprint(root.Val)
		dfsSer(root.Left, ret)
		dfsSer(root.Right, ret)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ret := ""
	dfsSer(root, &ret)
	return ret
}

func dfsDeser(data *string, i *int, pVal int) *TreeNode {
	if *i >= len(*data) {
		return nil
	}

	l := *i
	for (*data)[*i] != ',' {
		*i++
	}
	val, _ := strconv.ParseInt((*data)[l:*i], 10, 0)
	node := &TreeNode{
		Val: int(val),
	}
	if *i+1 < len(*data) {
		*i++
	}
	if *i+1 < len(*data) {
		r := *i
		for (*data)[r] != ',' {
			r++
		}
		nextVal, _ := strconv.ParseInt((*data)[*i:r], 10, 0)
		if nextVal < val {
			node.Left = dfsDeser(data, i, int(val))
		}
	}
	if *i+1 < len(*data) {
		r := *i
		for (*data)[r] != ',' {
			r++
		}
		val, _ := strconv.ParseInt((*data)[*i:r], 10, 0)
		if int(val) >= pVal {
			return node
		}
	}
	node.Right = dfsDeser(data, i, pVal)
	return node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	i := 0
	return dfsDeser(&data, &i, 10001)
}
