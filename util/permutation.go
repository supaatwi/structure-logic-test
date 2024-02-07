package util

func Permutation(s string) []string {
	var permutations []string

	var permute func(string, string)
	permute = func(str, prefix string) {
		if len(str) == 0 {
			permutations = append(permutations, prefix)
			return
		}

		usedChars := make(map[rune]bool)
		for i, char := range str {
			if !usedChars[char] {
				remaining := str[:i] + str[i+1:]
				permute(remaining, prefix+string(char))
				usedChars[char] = true
			}
		}
	}

	permute(s, "")
	return permutations

}
