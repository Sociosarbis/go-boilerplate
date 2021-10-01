package city

func destCity(paths [][]string) string {
	cityMap := map[string]bool{}
	for _, p := range paths {
		cityMap[p[0]] = true
	}
	for _, p := range paths {
		if _, exists := cityMap[p[1]]; !exists {
			return p[1]
		}
	}
	return ""
}
