package tracker

const n int = 1e5 + 1

type FrequencyTracker struct {
	numFreq   [n]int
	freqCount [n]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{}
}

func (this *FrequencyTracker) Add(number int) {
	oldFreq := this.numFreq[number]
	this.numFreq[number]++
	this.freqCount[oldFreq+1]++
	if oldFreq != 0 {
		this.freqCount[oldFreq]--
	}
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.numFreq[number] == 0 {
		return
	}
	oldFreq := this.numFreq[number]
	this.numFreq[number]--
	if oldFreq > 1 {
		this.freqCount[oldFreq-1]++
	}
	this.freqCount[oldFreq]--

}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freqCount[frequency] != 0
}
