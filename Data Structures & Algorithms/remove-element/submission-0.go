func removeElement(nums []int, val int) int {
    safePtr := 0

    for i, num := range nums {
        if num != val {
            nums[i], nums[safePtr] = nums[safePtr], nums[i]
            safePtr++
        }
    }

    return safePtr
}
