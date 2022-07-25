package inserter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root      *TreeNode
	prevLayer []*TreeNode
	count     int
}

func Constructor(root *TreeNode) CBTInserter {
	bfs := []*TreeNode{root}
	count := 0
	for {
		n := len(bfs)
		for i := 0; i < n; i++ {
			node := bfs[i]
			if node.Left != nil {
				bfs = append(bfs, node.Left)
			}
			if node.Right != nil {
				bfs = append(bfs, node.Right)
			}
		}
		if len(bfs) != n+(n<<1) {
			count = len(bfs) - n
			bfs = bfs[:n]
			break
		} else {
			bfs = bfs[n:]
		}
	}

	return CBTInserter{
		root,
		bfs,
		count,
	}
}

func (this *CBTInserter) Insert(val int) int {
	node := &TreeNode{
		val,
		nil,
		nil,
	}
	i := this.count >> 1
	if this.count&1 == 0 {
		this.prevLayer[i].Left = node
	} else {
		this.prevLayer[i].Right = node
	}

	ret := this.prevLayer[i].Val

	this.count++
	if this.count == len(this.prevLayer)<<1 {
		nextLayer := make([]*TreeNode, this.count)
		for i := 0; i < len(this.prevLayer); i++ {
			nextLayer[i<<1] = this.prevLayer[i].Left
			nextLayer[i<<1+1] = this.prevLayer[i].Right
		}
		this.prevLayer = nextLayer
		this.count = 0
	}
	return ret
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
