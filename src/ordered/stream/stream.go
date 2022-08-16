package stream

type OrderedStream struct {
	i      int
	values []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		i:      0,
		values: make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.values[idKey-1] = value
	j := this.i
	for j < len(this.values) && len(this.values[j]) != 0 {
		j++
	}
	ret := this.values[this.i:j]
	this.i = j
	return ret
}
