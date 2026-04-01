func getConcatenation(nums []int) []int {
    length := len(nums)
    ans := make([]int, 2 * length)

    for i := 0; i < 2*length; i++ {
        ans[i] = nums[i%len(nums)]
    }

    return ans
}
