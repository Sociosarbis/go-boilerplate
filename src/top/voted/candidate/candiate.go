package candidate

type TopVotedCandidate struct {
	times     *[]int
	topVoteds *[]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	votedNums := make([]int, len(persons))
	maxVal := 0
	for i := range persons {
		votedNums[i] = 0
	}
	topVoteds := make([]int, len(times))
	for i := range times {
		votedNums[persons[i]] += 1
		nextMax := votedNums[persons[i]]
		if nextMax >= maxVal {
			maxVal = nextMax
			topVoteds[i] = persons[i]
		} else {
			topVoteds[i] = topVoteds[i-1]
		}
	}
	return TopVotedCandidate{
		times:     &times,
		topVoteds: &topVoteds,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	l := 0
	r := len(*this.times) - 1
	for l <= r {
		mid := (l + r) >> 1
		if (*this.times)[mid] > t {
			r = mid - 1
		} else if (*this.times)[mid] < t {
			if mid+1 < len(*this.times) && (*this.times)[mid+1] <= t {
				l = mid + 1
			} else {
				l = mid
				break
			}
		} else {
			l = mid
			break
		}
	}
	return (*this.topVoteds)[l]
}
