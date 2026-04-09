func sortColors(nums []int) {
    zp := 0
    tp := len(nums) - 1
    i := 0

    for i <= tp {
        if nums[i] == 0 {
            nums[i], nums[zp] = nums[zp], nums[i]
            zp++
            i++
        } else if nums[i] == 2 {
            nums[i], nums[tp] = nums[tp], nums[i]
            tp--
            // do NOT increment i — re-examine the swapped-in element
        } else {
            i++ // nums[i] == 1, just advance
        }
    }
}