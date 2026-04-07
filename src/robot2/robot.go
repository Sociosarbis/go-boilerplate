package robot2

var dirs [4]string = [4]string{"East", "North", "West", "South"}

type Robot struct {
	width  int
	height int
	x      int
	y      int
	dir    int
}

func Constructor(width int, height int) Robot {
	return Robot{
		width:  width,
		height: height,
	}
}

func (this *Robot) Step(num int) {
	num %= (this.width + this.height - 2) * 2
	if num == 0 {
		if this.x == 0 {
			if this.y == 0 {
				this.dir = 3
			} else if this.y == this.height-1 {
				this.dir = 2
			}
		} else if this.x == this.width-1 {
			if this.y == 0 {
				this.dir = 0
			} else if this.y == this.height-1 {
				this.dir = 1
			}
		}
	}
	for num > 0 {
		var delta int
		var changeDir bool
		switch this.dir {
		case 0:
			changeDir = this.x == this.width-1
		case 1:
			changeDir = this.y == this.height-1
		case 2:
			changeDir = this.x == 0
		case 3:
			changeDir = this.y == 0
		}
		if changeDir {
			this.dir = (this.dir + 1) % 4
		}
		switch this.dir {
		case 0:
			if num >= this.width-1-this.x {
				delta = this.width - 1 - this.x
			} else {
				delta = num
			}
			this.x += delta
		case 1:
			if num >= this.height-1-this.y {
				delta = this.height - 1 - this.y
			} else {
				delta = num
			}
			this.y += delta
		case 2:
			if num >= this.x {
				delta = this.x
			} else {
				delta = num
			}
			this.x -= delta
		case 3:
			if num >= this.y {
				delta = this.y
			} else {
				delta = num
			}
			this.y -= delta
		}
		num -= delta
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	return dirs[this.dir]
}
