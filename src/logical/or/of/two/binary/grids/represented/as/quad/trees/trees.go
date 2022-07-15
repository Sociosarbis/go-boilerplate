package trees

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf && quadTree2.IsLeaf {
		quadTree1.Val = quadTree1.Val || quadTree2.Val
		return quadTree1
	}
	var ret *Node
	if !(quadTree1.IsLeaf || quadTree2.IsLeaf) {
		quadTree1.TopLeft = intersect(quadTree1.TopLeft, quadTree2.TopLeft)
		quadTree1.TopRight = intersect(quadTree1.TopRight, quadTree2.TopRight)
		quadTree1.BottomLeft = intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
		quadTree1.BottomRight = intersect(quadTree1.BottomRight, quadTree2.BottomRight)
		ret = quadTree1
	} else {
		if !quadTree1.IsLeaf {
			temp := quadTree2
			quadTree2 = quadTree1
			quadTree1 = temp
		}

		if quadTree1.Val {
			return quadTree1
		} else {
			return quadTree2
		}
	}
	if ret.TopLeft.IsLeaf && ret.TopRight.IsLeaf && ret.BottomLeft.IsLeaf && ret.BottomRight.IsLeaf {
		if ret.TopLeft.Val == ret.TopRight.Val && ret.TopRight.Val == ret.BottomLeft.Val && ret.BottomLeft.Val == ret.BottomRight.Val {
			ret.IsLeaf = true
			ret.Val = ret.TopLeft.Val
			ret.TopLeft = nil
			ret.TopRight = nil
			ret.BottomLeft = nil
			ret.BottomRight = nil
		}
	}
	return ret
}
