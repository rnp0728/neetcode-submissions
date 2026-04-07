func groupAnagrams(strs []string) [][]string {
	// store them in a map of [26]int -> []string
	// convert the strings to their []int value then store
	// into their respective list

    groups := make(map[[26]int][]string)

	for _, s := range strs {
		var count [26]int
		for i := 0; i < len(s); i++ {
			count[s[i] - 'a']++
		}
		groups[count] = append(groups[count], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
