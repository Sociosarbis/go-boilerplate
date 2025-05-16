package transformations2

import mat "github.com/boilerplate/src/matrix"

const mask int64 = 1e9 + 7

func lengthAfterTransformations(s string, t int, nums []int) int {
	cur := mat.NewMat(26, 26)
	for i := 0; i < 26; i++ {
		for j := 1; j <= nums[i]; j++ {
			cur.Set(i, (i+j)%26, 1)
		}
	}
	m := mat.NewIdentMat(26, 26)
	for t != 0 {
		if t&1 == 1 {
			m = m.Mul(&cur)
		}
		cur = cur.Mul(&cur)
		t >>= 1
	}
	var ret int64
	counter := [26]int{}
	for _, c := range s {
		index := c - 97
		counter[index]++
	}
	for i := 0; i < 26; i++ {
		if counter[i] != 0 {
			c := int64(counter[i])
			for j := 0; j < 26; j++ {
				ret = (ret + c*int64(m.Get(i, j))) % mask
			}
		}
	}
	return int(ret)
}
