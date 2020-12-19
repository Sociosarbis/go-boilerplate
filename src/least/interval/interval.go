package interval

import "sort"

type record struct {
	c byte
	i int
	n int
}

// Task Scheduler
// 优先执行数量最多的任务种类
func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 || n == 0 {
		return len(tasks)
	}
	records := []record{}
	counts := map[byte]int{}
	for _, t := range tasks {
		_, ok := counts[t]
		if !ok {
			counts[t] = 0
			records = append(records, record{t, -1, 0})
		}
		counts[t]++
	}

	for i, r := range records {
		records[i].n = counts[r.c]
	}

	sort.Slice(records, func(i, j int) bool {
		return records[j].n < records[i].n
	})
	ret := 0
	for len(records) != 0 {
		for i, r := range records {
			if ret-r.i > n || r.i == -1 {
				records[i].n--
				records[i].i = ret
				if records[i].n > 0 {
					if i+1 < len(records) {
						j := i + 1
						for ; j < len(records); j++ {
							if records[i].n >= records[j].n {
								break
							}
						}
						tmp := records[i]
						copy(records[i:j-1], records[i+1:j])
						records[j-1] = tmp
					}
				} else {
					copy(records[i:len(records)-1], records[i+1:])
					records = records[:len(records)-1]
				}
				break
			}
		}
		ret++
	}
	return ret
}
