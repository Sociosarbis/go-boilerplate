package hash_map

type Pair struct {
	k int
	v int
}

type MyHashMap struct {
	buckets [][]Pair
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		make([][]Pair, 1009),
	}
}

/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int) {
	i, j := m.FindIndex(key)
	if j == -1 {
		m.buckets[i] = append(m.buckets[i], Pair{key, value})
	} else {
		m.buckets[i][j].v = value
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
	i, j := m.FindIndex(key)
	if j != -1 {
		return m.buckets[i][j].v
	}
	return -1
}

func (m *MyHashMap) FindIndex(key int) (int, int) {
	i := key % 1009
	for j := 0; j < len(m.buckets[i]); j++ {
		if m.buckets[i][j].k == key {
			return i, j
		}
	}
	return i, -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int) {
	i, j := m.FindIndex(key)
	if j != -1 {
		m.buckets[i] = append(m.buckets[i][0:j], m.buckets[i][j+1:]...)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
