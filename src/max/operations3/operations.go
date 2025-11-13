package operations3

type chunk struct {
	next  *chunk
	end   int
	count int
}

func maxOperations(s string) int {
	var head, cur *chunk
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			if head == nil {
				head = &chunk{
					next:  nil,
					end:   i,
					count: 1,
				}
				cur = head
			} else {
				if cur.end+1 == i {
					cur.end = i
					cur.count++
				} else {
					next := &chunk{
						next:  nil,
						end:   i,
						count: 1,
					}
					cur.next, cur = next, next
				}
			}
		}
	}
	if head == nil {
		return 0
	}
	var out int
	for head.end+1 != n {
		if head.next != nil {
			out += head.count
			head.next.count += head.count
			head = head.next
		} else {
			out += head.count
			head.end = n - 1
		}
	}
	return out
}
