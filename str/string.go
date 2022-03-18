package str

func HighestFreq(str string) rune {
	countMap := make(map[rune]int)

	for _, c := range str {
		countMap[c]++
	}

	var mode rune
	maxCount := 0

	for char, count := range countMap {
		if count > maxCount {
			mode = char
			maxCount = count
		}
	}

	return mode
}
