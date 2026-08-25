// Edge Cases
// Only one element in array, return that element

// PseudoCode
// prefix is strs[0]
// compare the next element, 
// cut the prefix to the length of 
// from left to right with the longest prefix, the letter they diverge the prefix is the index before
// continue with all the elements
// return prefix

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func longestCommonPrefix(strs []string) string {
	longestPrefix := strs[0]
	for _, nextWord := range strs[1:] {
		
		shorterWordLength := min(len(nextWord), len(longestPrefix))
		longestPrefix = longestPrefix[:shorterWordLength]

		for i, letter := range nextWord[:shorterWordLength] {
			if rune(longestPrefix[i]) != letter {
				longestPrefix = longestPrefix[:i]
				break
			}
		}

	}
	return longestPrefix
}
