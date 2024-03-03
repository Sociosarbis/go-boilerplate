package stack

type MyStack struct {
	i     byte
	lists [2][]int
}

func Constructor() MyStack {
	return MyStack{
		0,
		[2][]int{},
	}
}

func (this *MyStack) Push(x int) {
	this.lists[this.i] = append(this.lists[this.i], x)
}

func (this *MyStack) Pop() int {
	n := len(this.lists[this.i])
	this.lists[1-this.i] = this.lists[this.i][0 : n-1]
	this.lists[this.i] = this.lists[this.i][n-1:]
	ret := this.lists[this.i][0]
	this.lists[this.i] = this.lists[this.i][:0]
	this.i = 1 - this.i
	return ret
}

func (this *MyStack) Top() int {
	n := len(this.lists[this.i])
	this.lists[1-this.i] = this.lists[this.i][0 : n-1]
	this.lists[this.i] = this.lists[this.i][n-1:]
	ret := this.lists[this.i][0]
	this.lists[1-this.i] = append(this.lists[1-this.i], ret)
	this.lists[this.i] = this.lists[this.i][:0]
	this.i = 1 - this.i
	return ret
}

func (this *MyStack) Empty() bool {
	return len(this.lists[0]) == 0 && len(this.lists[1]) == 0
}
