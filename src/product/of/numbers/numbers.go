package numbers

type ProductOfNumbers struct {
	prefix []int
	offset int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		prefix: make([]int, 0, 4*1e4),
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = append(this.prefix, 0)
		this.offset = len(this.prefix)
	} else {
		var prev int
		if len(this.prefix) != 0 {
			prev = this.prefix[len(this.prefix)-1]
		}
		if prev == 0 {
			this.prefix = append(this.prefix, num)
		} else {
			this.prefix = append(this.prefix, prev*num)
		}
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > len(this.prefix)-this.offset {
		return 0
	}
	if k == len(this.prefix)-this.offset {
		return this.prefix[len(this.prefix)-1]
	}
	return this.prefix[len(this.prefix)-1] / this.prefix[len(this.prefix)-k-1]
}
