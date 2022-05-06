package counter

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	lower := t - 3000
	for i := range this.requests {
		if this.requests[i] >= lower {
			if i > 0 {
				this.requests = this.requests[i:]
			}
			return len(this.requests)
		}
	}
	return 0
}
