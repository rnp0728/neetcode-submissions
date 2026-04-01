func groupAnagrams(strs []string) [][]string {
    anagramGroups := make(map[string][]string)

    for _, s := range strs {
        key := stringSort(s)
        anagramGroups[key] = append(anagramGroups[key], s)
    }

    result := make([][]string, 0)

    for _, group := range anagramGroups {
        result = append(result, group)
    }

    return result
}

func stringSort(s string) string {
    r := []rune(s)

    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })

    return string(r)
}

