package visits

import (
	"fmt"
	"strconv"
)

func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, s := range cpdomains {
		var count int
		for i := 0; i < len(s); i++ {
			if !(s[i] >= '0' && s[i] <= '9') {
				res, _ := strconv.ParseUint(s[0:i], 10, 32)
				count = int(res)
				break
			}
		}

		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '.' || s[i] == ' ' {
				domain := s[i+1:]
				m[domain] += count
				if s[i] == ' ' {
					break
				}
			}
		}
	}
	ret := make([]string, len(m))
	i := 0
	for s, c := range m {
		ret[i] = fmt.Sprintf("%v %s", c, s)
		i++
	}
	return ret
}
