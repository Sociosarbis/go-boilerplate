package characters

import "sort"

type Properties [][]int

func (this Properties) Len() int {
	return len(this)
}

func (this Properties) Less(i, j int) bool {
	if this[i][1] < this[j][1] {
		return true
	} else if this[i][1] > this[j][1] {
		return false
	} else {
		if this[i][0] < this[j][0] {
			return true
		} else {
			return false
		}
	}
}

func (this Properties) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func numberOfWeakCharacters(properties [][]int) int {
	sort.Sort(Properties(properties))
	attack := 0
	defense := 0
	nextAttack := 0
	ret := 0
	for i := len(properties) - 1; i >= 0; i-- {
		if properties[i][1] != defense {
			defense = properties[i][1]
			attack = nextAttack
			if properties[i][0] > nextAttack {
				nextAttack = properties[i][0]
			} else if properties[i][0] < attack {
				ret += 1
			}
		} else {
			if properties[i][0] < attack {
				ret += 1
			}
		}
	}
	return ret
}
