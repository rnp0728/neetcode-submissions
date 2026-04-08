func majorityElement(nums []int) int {
    n := len(nums)

	storage := make(map[int]int)

	for _, num := range nums {
		storage[num]++
	}

	for k, v := range storage {
		if v > n/2 {
			return k
		}
	}
	return -1
}
