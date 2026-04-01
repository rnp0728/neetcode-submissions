func twoSum(nums []int, target int) []int {
    store := make(map[int]int)

    for i, num := range nums {
        lookup := target - num
        if _, exists := store[lookup]; exists {
            return []int{store[lookup], i}
        }
        store[num] = i
    }

    return []int{}
}
