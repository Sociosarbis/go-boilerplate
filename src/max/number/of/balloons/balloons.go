package balloons

import (
	"github.com/boilerplate/src/utils"
)

func maxNumberOfBalloons(text string) int {
	counter := map[rune]int{
		97:  0,
		98:  0,
		108: 0,
		110: 0,
		111: 0,
	}
	for _, c := range text {
		if count, ok := counter[c]; ok {
			counter[c] = count + 1
		}
	}
	return utils.Min(counter[97], counter[98], counter[108]>>1, counter[110], counter[111]>>1)
}
