package files

import "sort"

type LetterLogs [][]string

func (this LetterLogs) Len() int {
	return len(this)
}

func (this LetterLogs) Less(i, j int) bool {
	k := 0
	l := 0
	for k < len(this[i][1]) || l < len(this[j][1]) {
		if k < len(this[i][1]) {
			if l < len(this[j][1]) {
				if this[i][1][k] < this[j][1][l] {
					return true
				} else if this[i][1][k] > this[j][1][l] {
					return false
				}
			}
		} else {
			return true
		}
		k++
		l++
	}

	k = 0
	l = 0

	for k < len(this[i][0]) || l < len(this[j][0]) {
		if k < len(this[i][0]) {
			if l < len(this[j][0]) {
				if this[i][0][k] < this[j][0][l] {
					return true
				} else if this[i][0][k] > this[j][0][l] {
					return false
				}
			}
		} else {
			return true
		}
		k++
		j++
	}

	return false
}

func (this LetterLogs) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func reorderLogFiles(logs []string) []string {
	digitLogs := []string{}
	letterLogs := LetterLogs([][]string{})
	i := 0
	for _, log := range logs {
		for i = 0; i < len(log); i++ {
			if log[i] == ' ' {
				break
			}
		}
		if log[i+1] >= 48 && log[i+1] <= 57 {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, []string{log[:i], log[i+1:]})
		}
	}

	sort.Sort(letterLogs)
	for i = 0; i < letterLogs.Len(); i++ {
		logs[i] = letterLogs[i][0] + " " + letterLogs[i][1]
	}
	for i = 0; i < len(digitLogs); i++ {
		logs[i+letterLogs.Len()] = digitLogs[i]
	}

	return logs
}
