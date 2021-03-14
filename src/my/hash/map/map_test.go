package hash_map

import (
	"testing"
)

type Suite struct {
	methods []string
	args    [][]int
	rets    []int
}

func TestMap(t *testing.T) {
	suite := Suite{
		[]string{"put", "put", "get", "get", "put", "get", "remove", "get"},
		[][]int{
			{1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2},
		},
		[]int{
			0, 0, 1, -1, 0, 1, 0, -1,
		},
	}

	hashMap := Constructor()

	for i := 0; i < len(suite.methods); i++ {
		if suite.methods[i] == "put" {
			hashMap.Put(suite.args[i][0], suite.args[i][1])
		} else if suite.methods[i] == "get" {
			ret := hashMap.Get(suite.args[i][0])
			if ret != suite.rets[i] {
				t.Errorf("Failed, expected %v but get %v", suite.rets[i], ret)
			}
		} else {
			hashMap.Remove(suite.args[i][0])
		}
	}
}
