package spanner

type item struct {
	i int
	p int
}

type StockSpanner struct {
	i     int
	items []item
}

func Constructor() StockSpanner {
	return StockSpanner{
		i:     0,
		items: []item{},
	}
}

func (this *StockSpanner) Next(price int) int {
	i := len(this.items) - 1
	for ; i >= 0; i-- {
		if this.items[i].p > price {
			break
		}
	}
	this.items = this.items[:i+1]
	var ret int
	if len(this.items) == 0 {
		ret = this.i + 1
	} else {
		ret = this.i - this.items[len(this.items)-1].i
	}
	this.items = append(this.items, item{
		i: this.i,
		p: price,
	})
	this.i++
	return ret
}
