package length

func dfs(str string, mark int, i int, temp string, counter *map[string]int) {
	if i < len(str) {
		temp = temp + str[i:i+1]
		if count, ok := (*counter)[temp]; ok {
			if count != -1 {
				if mark != count {
					(*counter)[temp] = -1
				} else {
					(*counter)[temp] = mark
				}
			}
		} else {
			(*counter)[temp] = mark
		}
		dfs(str, mark, i+1, temp, counter)
		dfs(str, mark, i+1, temp[:len(temp)-1], counter)
	}
}

func findLUSlength(strs []string) int {
	counter := map[string]int{}

	for i, s := range strs {
		dfs(s, i, 0, "", &counter)
	}

	ret := 0
	for s, count := range counter {
		if count != -1 {
			if len(s) > ret {
				ret = len(s)
			}
		}
	}
	if ret == 0 {
		return -1
	} else {
		return ret
	}
}
