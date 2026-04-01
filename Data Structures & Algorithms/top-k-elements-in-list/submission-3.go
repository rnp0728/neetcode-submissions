func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)

    for _, num := range nums {
        freqMap[num]++
    }

    type pair struct {
        num int
        freq int
    }
    pairs := make([]pair, 0, len(freqMap))
    
    for num, freq := range freqMap {
        pairs = append(pairs, pair{num: num, freq: freq})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].freq > pairs[j].freq
    })

    result := make([]int, k)

    for i := 0; i < k; i++ {
        result[i] = pairs[i].num
    }
    return result
}
