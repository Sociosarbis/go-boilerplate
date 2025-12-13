package coupons

import "sort"

type item struct {
	code     string
	category int
}

func categoryFromString(s string) int {
	switch s {
	case "electronics":
		return 0
	case "grocery":
		return 1
	case "pharmacy":
		return 2
	case "restaurant":
		return 3
	}
	return -1
}

func less(a, b string) bool {
	var n int
	if len(b) < len(a) {
		n = len(b)
	} else {
		n = len(a)
	}
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return len(a) < len(b)
}

func isValid(code string) bool {
	if len(code) == 0 {
		return false
	}
	for _, c := range code {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_') {
			return false
		}
	}
	return true
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	n := len(code)
	items := make([]item, 0, n)
	for i, c := range code {
		if isActive[i] && isValid(c) {
			category := categoryFromString(businessLine[i])
			if category != -1 {
				items = append(items, item{c, category})
			}
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].category < items[j].category || (items[i].category == items[j].category && less(items[i].code, items[j].code))
	})
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item.code)
	}
	return out
}
