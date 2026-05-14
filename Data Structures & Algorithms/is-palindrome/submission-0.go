func isPalindrome(s string) bool {
	// left and right pointer
	i, j := 0, len(s)-1

	for i < j {
		// move left pointer 
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}

		// move right pointer
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}