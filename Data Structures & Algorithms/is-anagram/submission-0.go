func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sentence map[rune]int = map[rune]int{}

	for _, c := range s {
		sentence[c]++
	}

	for _, c := range t {
		sentence[c]--
	}

	for _, v := range sentence {
		if v > 0 {
			return false
		}
	}

	return true
}
