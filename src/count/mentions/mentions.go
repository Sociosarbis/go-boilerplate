package mentions

import (
	"sort"
	"strings"
)

func parseInt(num string) int {
	var out int
	var base int = 1
	n := len(num)
	for i := n - 1; i >= 0; i-- {
		out += int(num[i]-'0') * base
		base *= 10
	}
	return out
}

type msg []int

func (self msg) time() int {
	return self[1]
}

func (self msg) isMsg() bool {
	return self[0] == 0
}

func (self msg) toHere() bool {
	return self[2] == -1
}

func (self msg) toAll() bool {
	return self[2] == -2
}

func (self msg) users() []int {
	return self[2:]
}

func fromEvent(event []string) msg {
	time := parseInt(event[1])
	if event[0] == "MESSAGE" {
		if event[2] == "HERE" {
			return msg{
				0,
				time,
				-1,
			}
		} else if event[2] == "ALL" {
			return msg{
				0,
				time,
				-2,
			}
		} else {
			msg := msg{0, time}
			for _, id := range strings.Split(event[2], " ") {
				msg = append(msg, parseInt(id[2:]))
			}
			return msg
		}
	} else {
		return msg{1, time, parseInt(event[2])}
	}
}

func countMentions(numberOfUsers int, events [][]string) []int {
	n := len(events)
	msgs := make([]msg, 0, n)
	for _, event := range events {
		msgs = append(msgs, fromEvent(event))
	}
	sort.Slice(msgs, func(i, j int) bool {
		return msgs[i].time() < msgs[j].time() || (msgs[i].time() == msgs[j].time() && !msgs[i].isMsg())
	})
	users := make([]int, numberOfUsers)
	out := make([]int, numberOfUsers)
	for _, msg := range msgs {
		time := msg.time()
		if msg.isMsg() {
			if msg.toHere() {
				for i := 0; i < numberOfUsers; i++ {
					if users[i] <= time {
						out[i]++
					}
				}
			} else if msg.toAll() {
				for i := 0; i < numberOfUsers; i++ {
					out[i]++
				}
			} else {
				for _, i := range msg.users() {
					out[i]++
				}
			}
		} else {
			for _, i := range msg.users() {
				users[i] = time + 60
			}
		}
	}
	return out
}
