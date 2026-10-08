func isPalindrome(s string) bool {
	ls := regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s, "")
	ls = strings.ToLower(ls)
	fmt.Println(ls)
	for i, _ := range ls {
		if ls[i] != ls[len(ls) - i - 1] {
			return false
		}
	}

	return true
}
