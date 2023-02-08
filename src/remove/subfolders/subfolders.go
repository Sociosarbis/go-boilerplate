package subfolders

import "strings"

type node struct {
	end  bool
	next map[string]*node
}

func collect(n *node, temp []byte, out *[]string) {
	if n.end {
		*out = append(*out, string(temp))
	} else {
		for k, next := range n.next {
			collect(next, append(temp, ("/"+k)...), out)
		}
	}
}

func removeSubfolders(folder []string) []string {
	root := &node{
		end:  false,
		next: map[string]*node{},
	}

	for _, f := range folder {
		n := root
		for _, seg := range strings.Split(f, "/")[1:] {
			if n.end {
				break
			}
			if nextN, ok := n.next[seg]; ok {
				n = nextN
			} else {
				nextNode := &node{
					end:  false,
					next: map[string]*node{},
				}
				n.next[seg] = nextNode
				n = nextNode
			}
		}
		n.end = true
	}
	ret := []string{}
	collect(root, []byte{}, &ret)
	return ret
}
