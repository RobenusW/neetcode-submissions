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
	for i := 1; i < len(strs); i++ {
		longestPrefix = longestPrefix[:min(len(strs[i]), len(longestPrefix))]
		for i, letter := range strs[i][:len(longestPrefix)] {
			if rune(longestPrefix[i]) != letter {
				longestPrefix = longestPrefix[:i]
				break
			}
		}

	}
	return longestPrefix
}
