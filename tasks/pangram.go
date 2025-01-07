package pangram

import "strings"

func pangram(input string) bool {
	s := strings.ToLower(input)

	for char := "a"; char <= "z"; char++ {
		if !strings.ContainsRune(s, char) {
			return false 
		}
	}

	return true
}
