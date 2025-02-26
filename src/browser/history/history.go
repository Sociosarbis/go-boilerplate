package history

type BrowserHistory struct {
	stack []string
	i     int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		stack: []string{homepage},
		i:     0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.stack = append(this.stack[:this.i+1], url)
	this.i++
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.i {
		steps = this.i
	}
	this.i -= steps
	return this.stack[this.i]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.i+steps >= len(this.stack) {
		steps = len(this.stack) - 1 - this.i
	}
	this.i += steps
	return this.stack[this.i]
}
