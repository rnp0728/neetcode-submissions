func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i, num := range nums {
		rem := target - num
		if val, ok := m[rem]; ok {
			return []int{val, i}
		}
		m[num] = i
	}
	return []int{}
}
